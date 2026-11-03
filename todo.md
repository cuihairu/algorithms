# algorithms 需求台账

> 验收口径：各模块 `go test ./...` 全绿；docs 站 `pnpm run docs:build` 成功；README/docs 与源码一致（见 docs/审计-文档一致性.md）。

## 已完成

- [x] array 泛型数组封装（越界检查），4 测试，证据 ede7480
- [x] sorts 冒泡/插入/选择/归并/快排/计数/桶 7 种排序，7 测试，证据 1df32b9
- [x] stack 接口 + 数组栈/链表栈，6 测试，证据 040b66c
- [x] linkedlist 头哨兵单链表 + 读写锁 + 快慢指针，11 测试（含 -race），证据 f7a6b7a
- [x] queue 接口 + 数组队列（头游标）/链表队列，6 测试，证据 c231579
- [x] tree 泛型 BST（插入/查/删/最值/中序/高度/平衡），8 测试，证据 1eb38b0
- [x] heap 泛型最小堆（上浮/下沉/O(n) 建堆），6 测试，证据 a64bf41
- [x] hash 链地址哈希表（FNV-1a + 负载因子倍增），5 测试，证据 d7a45a6
- [x] graph 邻接表有向/无向图 + BFS/DFS/最短路，9 测试，证据 1eb38b0
- [x] go.work 收录全部 9 个模块，证据 3c39425
- [x] docs/notes 每模块一篇笔记（API 表、复杂度、历史坑位），证据 b950a7f
- [x] docs 站双语化（en 根 + /zh/，README 双语 + 首屏切换器），证据 7d6b72c
- [x] docs 站迁移 VitePress（config.mts + zh locale + index.md 首页），证据 b572cd5

## 待办

- [x] sorts 冒泡/计数/插入/快排测试补断言（现仅 t.Log 烟测，曾掩盖 MergeSort P0；selection_test.go 为空文件），证据 sorts/go test 全绿 12 测试（含 Correctness 断言）
- [x] graph 无向 AddEdge(A,A) 自环向自身邻接表追加两次，EdgeCount 记 1（graph/go/graph.go:31-35，实证 Neighbors(a)=[a a]），证据 go test graph 全绿 10 测试（TestGraphUndirectedSelfLoop 回归）
- [x] linkedlist RemoveLast 冗余 isEmpty 判断（死代码），证据 go test linkedlist 全绿 11 测试
- [x] queue 补编译期接口守卫（stack 已加，queue 未加），证据 go test queue 全绿 6 测试
- [x] 说明非线程安全：仅 linkedlist 文档化 RWLock，queue/stack/tree/heap/hash/graph 未声明并发限制，证据 stack/queue/tree/heap/hash/graph 6 模块包注释已声明 + 全绿
- [x] hash 桶掩码整数转换 gosec G115 告警（hash/go/hashmap.go:104；桶数恒为 2 的幂无实际溢出，误报，加注释或忽略标记），证据 #nosec G115 + gosec Issues:0 + go test hash 全绿
- [x] array Insert 先判 FullArray 后判越界：满数组越界插入返回 full array 而非 out of range（array/go/array.go:50-53），证据 go test array 全绿 6 测试（TestFullArrayOutOfRangeInsert 回归）
- [x] linkedlist Insert 对 index>size+1 静默尾插、index<=0 静默返回，无越界反馈（linkedlist/go/linkedlist.go:173-188），证据 go test linkedlist 全绿 11 测试（ErrIndexOutOfRange 回归）
- [x] ArrayQueue.Enqueue 恒定 append，非空稳态下内存 O(总操作数) 而非 O(当前大小)（queue/go/array.go:27-30），证据 go test queue 全绿 7 测试（TestArrayQueueBoundedMemory 回归，容量有界）
- [x] docs/demo、docs/guide、docs/slides 为旧 VuePress 模板页（@include 等特性 VitePress 不支持），迁移或删除，证据删除模板页 + config/index 清理 + pnpm run docs:build 成功 45.3s
- [x] 清理旧 VuePress 依赖（vuepress/vuepress-theme-hope/@vuepress/client），证据 package.json 仅剩 vue+vitepress，lockfile 无 vuepress 引用，pnpm run docs:build 成功 14.2s
- [ ] 启用 GitHub Pages 并验证徽章可达：仓库私有 + Free 套餐不支持（422），Actions 未产生部署 run，需公开仓库或升级套餐（部署 workflow 已改为 VitePress + pnpm） — BLOCKED：缺用户决策（公开仓库或升级付费套餐），workflow 已就绪，无法在本环境验证

## 审核发现

- [x] 审核发现：报告 B6 行号引用 linkedlist/go/linkedlist.go:105-108 指向 `Reverse()`，实际重复 isEmpty 判断在 `RemoveLast` 的 140-145（死代码在 143-145）；行号需修正（docs/review/2026-11-02-code-review-opencode.md:34），证据报告已改为 141-146 并标 **已修** `0c24336`
- [x] 审核发现：报告 B10 行号引用 array/go/array.go:32-43 起点错误，`Insert` 实际在 49-55（先判 FullArray 在 50-52，越界判在 53-54）；行号需修正（docs/review/2026-11-02-code-review-opencode.md:38），证据报告已改为 49-55（越界 50-52/满 53-55）并标 **已修** `819c537`
- [x] 审核发现：报告 B5 行号 graph/go/graph.go:24-37 覆盖 AddEdge 整体（24-36），但 todo.md 对应条目标 31-35；两处行号口径不一致，需统一（docs/review/2026-11-02-code-review-opencode.md:33），证据报告/todo 已统一为 27-40 并标 **已修** `a089271`
- [x] 审核发现：报告 B2 行号 sorts/go/merge.go:20-23 指向当前已修复代码，当前 merge 的左半余量回拷在 21-24、右半余量在 25-28；报告中「20-23」为历史缺陷行号，对已修复代码需标注 commit 前行号或改为当前文件行号（docs/review/2026-11-02-code-review-opencode.md:30），证据报告已标注历史行号 20-23 + 当前 21-28 并保留 **已修** `9edeae7`

> 独立审核（3eed44c361d6..HEAD，2026-11-02）：逐文件复核该区间 16 个提交（Go 修复+测试、docs 站迁移/依赖清理、go.work/lock），9 模块 `go test ./...` 与 `linkedlist/queue -race` 复跑全绿；代码侧未见新的越界/泄漏/竞态/注入/反序列化缺陷，发现集中在文档漂移与测试缺口。

- [x] 审核发现：`LinkedList.Insert` 返回 `error` 为破坏性公开 API 变更，但 docs/zh/notes/linkedlist.md:26 与 docs/notes/linkedlist.md:26 仍写「index<=0 不动作；越界等价尾插」，与 `321c223` 后行为（越界返回 ErrIndexOutOfRange、`index<=0` 同样报错）矛盾且未记载新返回类型，违反验收口径「README/docs 与源码一致」（linkedlist/go/linkedlist.go:174），证据 docs/notes/linkedlist.md:26 与 docs/zh/notes/linkedlist.md:26 已改为合法区间 `[1,size+1]` + 返回 error/越界 ErrIndexOutOfRange
- [x] 审核发现：`ArrayQueue.Enqueue` 压实修复后，docs/zh/notes/queue.md:39 与 docs/notes/queue.md:39 仍称「非空期间底层切片随 Enqueue 持续增长（只在清空时回收）」，与 `7777716` 的 `head>=len/2` 前移压实（内存有界）矛盾，笔记陈旧（queue/go/array.go:35-39），证据 docs/notes/queue.md:39 与 docs/zh/notes/queue.md:39 已改为「已消费至少一半时前移压实，内存有界」
- [x] 审核发现：TestArrayQueueBoundedMemory 仅断言 `Size()` 与 `cap`，未断言压实后出队元素仍保持 FIFO 顺序，若压实逻辑错位（丢元素/乱序）测试无法捕获（queue/go/array_test.go:96-111），证据 queue/go/array_test.go:112-120 已断言稳态末尾 16 元素 FIFO 顺序（9984..9999），queue go test 全绿
- [x] 审核发现：TestLinkedListInsert 未覆盖 `index==1` 头插与空链表合法插入（`size+1==1`）边界，越界/合法边界覆盖不完整（linkedlist/go/linkedlist_test.go:123-156），证据 linkedlist/go/linkedlist_test.go:157-189 新增 TestLinkedListInsertBoundaries（index==1 头插/空链表合法插入/空表越界不修改），linkedlist go test -race 全绿
- [x] 审核发现：新增排序 Correctness 用例值域过小，未覆盖 `CountingSort` 在 `{math.MinInt, math.MaxInt}` 时 `max-min+1` 溢出为 0、`count[arr[i]-min]` 越界 panic 的风险（实现未改，但新测试面向该函数）（sorts/go/counting_test.go:18-40），证据 sorts/go/counting.go:25-29 值域溢出时退化为比较排序 + sorts/go/counting_test.go:43-62 新增 TestCountingSortOverflowRange（含 math.MinInt/MaxInt 对），sorts go test 全绿 + gosec Issues:0

> 审核任务（只审不改，16438291cdba..HEAD，2026-11-03）：逐文件复核 7 个提交（counting sort 溢出守卫、linkedlist/queue 边界测试、docs 对齐、todo 记账）。已复跑 sorts / linkedlist / queue（含 -race）全部 PASS（TestCountingSortOverflowRange -count=1 亦绿）；代码侧未发现注入/反序列化/竞态/错误处理缺失，发现集中在 counting sort 守卫的相邻崩溃路径与其测试缺口。

- [x] 审核发现：`CountingSort` 新增溢出守卫只覆盖 `max-min+1<=0`（回绕为 0/负）；值域巨大但未回绕（如 `{0,1<<40}`、`{0,math.MaxInt/2}`）仍通过守卫进入 `make([]int,max-min+1)` 尝试分配数 TB 级计数数组，随后 `for i:=1; i<=int(max-min); i++` 前缀和循环按 O(值域) 次数读写该数组 → 运行时 OOM / 近似挂起（makeslice 超大或 overcommit 后触达 8TiB）。这与本提交「防 range overflow panic」的意图属同一类崩溃，仅修复了回绕分支、遗漏了相邻未回绕大值域分支（sorts/go/counting.go:26-31），证据 sorts/go/counting.go:25-29 改用 uint64 跨度 `span>=1<<20` 阈值（超阈值退化稳定比较排序，杜绝 TB 级 make）+ sorts/go/counting_test.go:64-94 新增 uint 大值域 `{0,1<<40}` 等用例，sorts go test 全绿

- [x] 审核发现：溢出守卫退化分支用 `sort.Slice`（sorts/go/counting.go:27），为不稳定排序；而主路径（counting.go:39-44 逆序前缀和放置）为稳定计数排序。值域溢出时稳定性语义被悄然改变——当前 docs 未声明 counting sort 稳定，故非契约违反，但属隐式行为漂移，未来若调用方/文档依赖稳定性将出错（sorts/go/counting.go:27）；另 int8/int16 大值域（如 `{[-128,50]}`，真实值域 179 桶本可容纳）经类型 T 的 `max-min+1` 提前回绕为 ≤0，同样落入比较排序退化，属计数排序对窄整型的性能降级（sorts/go/counting.go:26），证据 sorts/go/counting.go:27 改用 `sort.SliceStable`（保持稳定语义，docs/notes/sorts.md:26 已声明稳定）+ 计数索引统一经 uint64 无回绕（int8 大值域走计数主路径），sorts go test 全绿

- [x] 审核发现：`TestCountingSortOverflowRange` 仅实例化 `int`，未覆盖两条相邻路径——无符号回绕（`CountingSort[uint]{0, math.MaxUint}`，`MaxUint+1` 回绕为 0）与大值域未回绕（`{0,1<<40}`）；后者现仍会 OOM/挂起，故现有测试无法捕获上一缺陷（sorts/go/counting_test.go:43-62），建议补 uint 与大值域用例并在守卫中改用「值域超过阈值即退化比较排序」，证据 sorts/go/counting_test.go:64-94 新增 TestCountingSortOverflowTypedRanges（int8 MinInt8/MaxInt8 对 + uint MaxUint 回绕对 + uint {0,1<<40} 大值域），sorts go test 全绿

> 审核任务（只审不改，a71f7c6ed25d15f10527303169e14adb9ddb4c70..HEAD，2026-11-03）：逐文件复核 1 个提交 0db777d（counting sort 溢出守卫改造 + 新增窄整型/无符号大值域测试 + todo 记账）。已复跑 sorts go test -race 全绿；GOARCH=386 go build ./... 通过（主实现经 uint64 跨度无 32 位问题），但 GOARCH=386 go test 编译失败。主路径正确性已人工推演：span=uint64(max)-uint64(min) 对有/无符号、跨 0、窄整型均无回绕，span<1<<20 时索引 ∈[0,span] 且 span+1≤1<<20 可安全 int 化，前缀和/逆序放置无越界；阈值边界 span==1<<20-1 与 1<<20 均无 off-by-one。未发现注入/反序列化/越界/泄漏/竞态/错误处理缺失。

- [x] 审核发现：新增 TestCountingSortOverflowTypedRanges 在 `[]uint` 字面量写入 `1 << 40`（1099511627776），超出 32 位平台 uint 上限（2^32-1），GOARCH=386/arm 下 `go test`/`go vet` 报 `cannot use 1 << 40 ... overflows uint` 编译失败，使 sorts 包在 32 位平台无法构建（sorts/go/counting_test.go:82）。主实现无此问题（GOARCH=386 go build ./... 通过），仅新测试用例破坏可移植性；建议改为平台可表示值或对 64 位分支单独断言。，证据 sorts/go/counting_test.go:82 改为 `1 << 31`（32 位可表示）+ sorts/go/counting_test.go:63-93 保持其余用例，sorts go test + GOARCH=386 go test 全绿

- [x] 审核发现：阈值边界未覆盖——现有用例值域要么 ≤255 要么远超 1<<20，均不落在 `span == 1<<20-1`（计数路径最大跨度）与 `span == 1<<20`（首次退化）边界，若比较符方向或 `make([]int, span+1)` 边界发生 off-by-one 回归将无法被捕获（sorts/go/counting_test.go:63-93 覆盖范围，sorts/go/counting.go:29-34 边界逻辑）；属次要测试缺口，主实现当前边界行为经人工推演正确。，证据 sorts/go/counting_test.go:94-106 新增 TestCountingSortRangeThresholds（`{1<<20-1,0}` 计数最大跨度 + `{1<<20,0}` 首次退化），sorts go test 全绿
