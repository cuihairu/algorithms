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
