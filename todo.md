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

- [ ] sorts selection 排序补测试（selection_test.go 为空文件）
- [ ] array Insert/Pop 允许 index>len 造成空洞（零值写入）
- [ ] linkedlist RemoveLast 冗余 isEmpty 判断
- [ ] queue 补编译期接口守卫（stack 已加，queue 未加）
- [ ] docs/demo、docs/guide、docs/slides 为旧 VuePress 模板页（@include 等特性 VitePress 不支持），迁移或删除
- [ ] 清理旧 VuePress 依赖（vuepress/vuepress-theme-hope/@vuepress/client）
- [ ] 启用 GitHub Pages 并验证徽章可达：仓库私有 + Free 套餐不支持（422），Actions 未产生部署 run，需公开仓库或升级套餐（部署 workflow 已改为 VitePress + pnpm）
