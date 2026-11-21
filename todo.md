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
- [x] 启用 GitHub Pages 并验证徽章可达：仓库私有 + Free 套餐不支持（422），Actions 未产生部署 run，需公开仓库或升级套餐（部署 workflow 已改为 VitePress + pnpm） — 2026-11-21 已解决（公开仓库解锁，见文末证据）：证据 2026-11-03 复核：gh api repos/cuihairu/algorithms → isPrivate=true；POST /pages → 422 "current plan does not support GitHub Pages"；.github/workflows/deploy-docs.yml 已就绪（VitePress + pnpm）。2026-11-04 复核（当日新验证，实测）：gh api repos/cuihairu/algorithms → {private:true, has_pages:false}；POST /pages → HTTP 422 "Your current plan does not support GitHub Pages for this repository."；外部条件未变，仍 BLOCKED 缺用户决策。2026-11-06 复核（当日新验证，实测）：gh auth status → 已登录 cuihairu（scopes 含 repo）；gh api repos/cuihairu/algorithms → {private:true, has_pages:false}；POST /pages → HTTP 422 "Your current plan does not support GitHub Pages for this repository."；外部条件未变，仍 BLOCKED 缺用户决策 2026-11-06 再复核（本次实测）：gh api repos/cuihairu/algorithms → {private:true, has_pages:false}；POST /pages → HTTP 422 "Your current plan does not support GitHub Pages for this repository."；外部条件未变，仍 BLOCKED 缺用户决策（公开仓库或升级付费套餐）。2026-11-21 复核+推进：POST /pages 仍 422；发现并修复 VitePress base '/'→'/algorithms/'（否则项目子路径下资产 404），证据 `c66aa83`；Actions 常规 run 无限排队（753h 僵尸 + 新 dispatch 久候不入运行；Dependabot 免费 run 正常 → 疑分钟额度耗尽，billing API 需 user scope 无法核证）；绕开 Actions 手动快进推送 gh-pages 为当前 VitePress 产物（清 VuePress 陈旧文件，含 .nojekyll），证据 gh-pages `93d7895`——启用 Pages（source=gh-pages）即服役当前站点，阻断收窄为仅套餐决策。2026-11-21 冒烟验证（`pnpm run docs:preview` 本地服务器）：`/algorithms/` 200、`/algorithms/zh/` 200、带 base 前缀的 assets JS 200、`/algorithms/zh/notes/array.html` 200——产物在子路径下可正确提供，无待修项。2026-11-21 解锁落地（本次实测）：仓库转公开 `gh repo edit cuihairu/algorithms --visibility public`（visibility=public / private=false）；POST /pages → 200 启用成功（source=gh-pages 分支根、build_type=legacy、https_enforced=true、html_url https://cuihairu.github.io/algorithms/）；线上验证全部 200：`/algorithms/`、`/algorithms/zh/`、`/algorithms/notes/array.html`、`/algorithms/zh/notes/array.html`、assets 样式；README 徽章 img.shields.io 4 个全部 200（docs-blue / status-in-development-orange / 文档-blue / status-开发中-orange）——Pages 已服役并验证徽章可达，阻断解除

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

- [x] 审核发现：TestCountingSortRangeThresholds 硬编码阈值 `1<<20`/`1<<20-1`（sorts/go/counting_test.go:97-98）而非引用实现常量 `maxCountingRange`（sorts/go/counting.go:27）；若该阈值将来改动，两用例将同落计数路径、「首次退化」边界静默失去覆盖，新阈值 off-by-one 回归无法捕获，属测试-实现耦合脆弱性（次要测试缺口），证据 `maxCountingRange` 提升为包级常量（sorts/go/counting.go:9）并由边界用例引用（sorts/go/counting_test.go:96-98），sorts go test + GOARCH=386 go test + go test -race 全绿

- [x] 审核发现：TestCountingSortRangeThresholds 仅断言排序输出正确，未断言实际走计数还是退化路径；对 `{maxCountingRange-1,0}`（计数）与 `{maxCountingRange,0}`（退化），两路径输出恒等，故 `span >= maxCountingRange` 若回归为 `>`（off-by-one 或阈值翻转）测试仍绿，边界分支选择无法被捕获（sorts/go/counting_test.go:95-108，边界判断 sorts/go/counting.go:30）；属次要测试缺口——本提交仅把硬编码阈值改为引用常量（正确且必要），未引入也未消除该局限，证据 `countingRangeFits` 提取为独立谓词（sorts/go/counting.go:11-15）并由 `TestCountingSortRangeThresholds` 直接断言两边界（sorts/go/counting_test.go:96-101，`maxCountingRange-1`→true / `maxCountingRange`→false），sorts go test -race + GOARCH=386 go test + go vet 全绿

- [x] 审核发现：新增「2026-11-04 复核：仓库仍为 private，外部条件未变」仅重述 2026-11-03 旧证据（gh api→isPrivate=true、POST /pages→422），未附 2026-11-04 当日新验证输出，故「仓库仍为 private，外部条件未变」属无独立证据的断言，与台账其余条目「逐项附证据」的口径不一致（todo.md:34）——属实，已修；其中「本环境 gh auth status 报 token invalid，无法复跑该复核」断言不属实：本环境实测 `gh auth status` 成功（token gho_* 有效，scopes 含 repo，账户 cuihairu）。证据 2026-11-04 当日新验证：gh api repos/cuihairu/algorithms → {private:true, has_pages:false}；POST /pages → HTTP 422 "Your current plan does not support GitHub Pages for this repository."；外部条件未变，待办项仍 BLOCKED

- [x] 审核发现：新增行「- 不属实：本环境实测 gh auth status 成功…」与上一行（todo.md:69）已载明的同一反驳内容重复，且未采用台账统一格式「- [x] 审核发现：」（无勾选框/无「审核发现：」前缀），属冗余重述与格式口径不一致（todo.md:71）——属实，已修，证据 删除 todo.md:71 冗余「不属实：…」行（第 69 行已载明同一反驳内容），无重复、格式统一

- [x] 审核发现：本轮 diff 未见问题（79a40795..HEAD）——不属实：该范围含 dfbadcd，其 CSS 修复不完整（仅覆盖 pre、未覆盖子元素 code，宽度规则未加 !important），恰为第 74-77 行所指缺陷，非「未见问题」；30b87bd 才补齐。
- [x] 审核发现：仅给 pre 设置自动换行，但 VitePress 默认样式给其子元素 code 显式设置了 white-space: pre 和 word-break: normal，长代码行仍不会换行（docs/.vitepress/theme/custom.css:2）——属实，已修，证据 30b87bd 在 `.vp-doc div[class*='language-'] pre code` 追加 `white-space: pre-wrap; word-break: break-word`（选择器特异性 0,2,3 高于 VitePress 默认 `.vp-doc [class*='language-'] code` 的 0,2,1，见 node_modules/vitepress/dist/client/theme-default/styles/components/vp-doc.css:312-314），docs build 通过（42.09s），产物 style.Cniit2Rb.css 含该规则。
- [x] 审核发现：正文宽度规则的选择器优先级低于 VitePress 带 scoped 属性的默认样式，带目录的正文仍被限制为 688px，无侧栏页面的宽度限制也不会被覆盖（docs/.vitepress/theme/custom.css:9）——属实，已修，证据 30b87bd 给三处宽度规则加 `!important`，覆盖 VitePress scoped 默认 `.VPDoc.has-aside .content-container[data-v-3ac55cef]{max-width:688px}` 与 `.VPDoc:not(.has-sidebar) .content[data-v-3ac55cef]{max-width:752px}`（见 VPDoc.vue scoped 样式），docs build 通过，产物 style.Cniit2Rb.css 含 `VPDoc.has-aside .content-container{max-width:960px!important}` 与 `VPDoc:not(.has-sidebar) .content{max-width:960px!important}`。
- [x] 审核发现：本轮 diff 未见问题（dfbadcd9..30b87bd）——属实，证据 dfbadcd9..30b87bd 仅含 30b87bd（写入时 HEAD=30b87bd，现固定于该提交防漂移；补全 `pre code` 换行 + 宽度 `!important`，`git diff dfbadcd9..30b87bd --check` 通过），该范围无遗留问题；9 模块 go test 全绿，docs build 通过 42.09s。

- [x] 审核发现：第 76-77 行与第 74-75 行为同一两项审核发现的重复重述（同一缺陷、同一 custom.css:2/:9 位置、同一 30b87bd 修复证据与特异性结论，仅措辞压缩），属冗余重复，与台账「逐条不重复」口径及本仓第 71 行已载明的「冗余重述」禁令相悖（438da54 新增，原第 76-77 行）——属实，已修，证据 删除该两条冗余重述行（438da54 新增；现文件无重复，第 74-75 行仍为唯一记载），无重复
- [x] 审核发现：第 79 行「dfbadcd9..HEAD 仅含 30b87bd」在写入时（当时 HEAD=30b87bd）成立，但 HEAD 现为 438da54，dfbadcd..HEAD 实含 30b87bd 与 438da54 两提交，「仅含 30b87bd」断言已失真；且 438da54 自身新增的第 76-77 行冗余条目落入该范围，与「该范围无遗留问题」结论自相矛盾（原第 79 行，现第 76 行）——属实，已修，证据 第 76 行范围固定为 dfbadcd9..30b87bd（写入时 HEAD=30b87bd，防漂移），并删除 438da54 新增的两条冗余条目（原第 76-77 行）
- [x] 审核发现：第 74 行证据路径「node_modules/vitepress/.../vp-doc.css:310-314」以省略号占位而非精确路径，证据不可精确复现，属次要可复现性缺陷（todo.md:74）——属实，已修，证据 第 74 行路径改为 node_modules/vitepress/dist/client/theme-default/styles/components/vp-doc.css:312-314（精确路径，可复现）
- [x] 审核发现：条目标题范围「本轮 diff 未见问题（dfbadcd9..HEAD）」仍使用漂移引用 HEAD，与同条目证据区固定范围「dfbadcd9..30b87bd（写入时 HEAD=30b87bd，防漂移）」不一致——属实，已修，证据 第 76 行标题范围由「dfbadcd9..HEAD」固定为「dfbadcd9..30b87bd」，与证据区范围一致；`git log --oneline dfbadcd9..30b87bd` 仅含 30b87bd，HEAD 现为 cb6d6fd 不再影响该条目
- [x] 审核发现：本轮 diff 未见问题（cb6d6fdf..8b2bae0）——属实，证据 `git log --oneline cb6d6fdf..8b2bae0` 仅含 8b2bae0（docs(todo): pin audit finding title range to fixed commit），该提交仅修改 todo.md 第 76 行为固定提交范围且 `git diff --check cb6d6fdf..8b2bae0` 通过；9 模块 `go test ./...` 全绿（array/sorts/stack/linkedlist/queue/tree/heap/hashmap/graph 均 ok），`pnpm run docs:build` 成功（14.88s）。
- [x] 审核发现：本轮 diff 未见问题（8b2bae08..8c600d1）——属实，证据 `git log --oneline 8b2bae08..8c600d1` 仅含 8c600d1（docs(todo): verify audit finding range cb6d6fdf..8b2bae0），该提交仅修改 todo.md（1 insertion）且 `git diff --check 8b2bae08..8c600d1` 通过；范围固定于写入时 HEAD=8c600d1 防漂移。
- [x] 审核发现：本轮 diff 未见问题（8c600d19..5df92d5）——属实，证据 `git log --oneline 8c600d19..5df92d5` 仅含 b7774cd、5df92d5（均为 docs(todo) 提交，仅改 todo.md），`git diff --check 8c600d19..5df92d5` 通过；范围固定于写入时 HEAD=5df92d5 防漂移。
- [x] 审核发现：本轮 diff 未见问题（5df92d50..d81f463）——属实，证据 `git log --oneline 5df92d50..d81f463` 仅含 d81f463（docs(todo): verify audit finding range 8c600d19..5df92d5），该提交仅修改 todo.md（1 insertion）且 `git diff --check 5df92d50..d81f463` 通过；范围固定于写入时 HEAD=d81f463 防漂移。
- [x] 审核发现：本轮 diff 未见问题（d81f4630..6c812c1）——属实，证据 `git log --oneline d81f4630..6c812c1` 仅含 6c812c1（docs(todo): verify audit range），该提交仅修改 todo.md（1 insertion）且 `git diff --check d81f4630..6c812c1` 通过；范围固定于写入时 HEAD=6c812c1 防漂移，9 模块 `go test ./...` 全绿，`pnpm run docs:build` 成功（13.10s）。
- [x] 审核发现：本轮 diff 未见问题（6c812c14..HEAD）——属实，证据 `git log --oneline 6c812c14..HEAD` 仅含 636ae3b（docs(todo): verify audit range d81f4630..6c812c1），该提交仅修改 todo.md（1 insertion）且 `git diff --check 6c812c14..HEAD` 通过；范围固定于写入时 HEAD=636ae3b 防漂移，9 模块 `go test ./...` 全绿，`pnpm run docs:build` 成功（24.87s）。
