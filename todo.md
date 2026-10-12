# algorithms 巡检台账（patrol-20261013-C1 建账）

盘点基准：4 个 Go module（array / sorts / stack / linkedlist）+ vuepress docs 站。
统一验收口径：各 module 内 `go test ./...` 全绿；docs 项以 `npm run docs:build` 成功 + 内容非模板为准。

## 一、已实现（代码 + 测试通过）

- [x] array/go 数组封装，含 4 个测试函数，`go test ./...` ok —— 证据 ede7480
- [x] sorts/go 冒泡/插入/选择/归并/快排/计数 6 种排序，各有测试且通过 —— 证据 934cb3b、71a2c03、1df32b9
- [x] linkedlist/go 双向链表 18 个方法实现完整 —— 证据 655b3bc（注意：测试为空壳，见二档）

## 二、部分实现（代码在，但有 bug 或测试缺失）

- [ ] sorts/go 桶排序 TestBucketSort panic：bucket.go:32 桶序号误用 nBuckets 整除（应为桶宽），全负输入 [-10..0] 时 index 越界。历史上 14d87d4 曾"fix bucket sort"但仍红。
  验收：sorts 模块 `go test ./...` 全绿，含全负用例。期次：C1（本轮修复）
- [ ] go.work 只收录 array/sorts，stack、linkedlist 两 module 游离在工作区外，模块内 `go test ./...` 报 "does not contain modules listed in go.work"
  验收：go.work 收录全部 4 个 module，四个模块内测试命令均可执行。期次：C2
- [ ] stack/go ArrayStack 全部方法已实现（Push/Pop/Top/Size/Flush/IsEmpty/String）但零测试 —— 证据 dc24bd1
  验收：补齐接口各方法用例并全绿。期次：C2
- [ ] stack/go LinkedListStack 仅 NewLinkedListStack + IsEmpty，Push/Pop/Top/Size/Flush 未实现
  验收：实现 Stack 接口全部方法，测试全绿。期次：C2
- [ ] linkedlist/go linkedlist_test.go 空壳（TestLinkedList 无断言）
  验收：覆盖 AddFirst/AddLast/RemoveFirst/RemoveLast/Insert/Reverse/HasCycle/FindMiddleNode。期次：C3

## 三、只有设计 / 完全没有

- [ ] docs/ 全部为 vuepress-theme-hope 模板原样页（docs/README.md 证据 c18aea3），无任何真实算法笔记 —— 模板不记已完成
  验收：为 array/sorts/stack/linkedlist 各写至少一篇笔记，`npm run docs:build` 成功。期次：C3
- [ ] queue / tree / graph / heap / hash 等结构完全没有目录
  验收：按「module 骨架 + 基础实现 + 测试」推进，一个结构一个 commit。期次：C4 起
