# algorithms 巡检台账（patrol-20261013-C1 建账）

盘点基准：4 个 Go module（array / sorts / stack / linkedlist）+ vuepress docs 站。
统一验收口径：各 module 内 `go test ./...` 全绿；docs 项以 `npm run docs:build` 成功 + 内容非模板为准。

## 一、已实现（代码 + 测试通过）

- [x] array/go 数组封装，含 4 个测试函数，`go test ./...` ok —— 证据 ede7480
- [x] sorts/go 冒泡/插入/选择/归并/快排/计数 6 种排序，各有测试且通过 —— 证据 934cb3b、71a2c03、1df32b9
- [x] linkedlist/go 双向链表 18 个方法实现完整 —— 证据 655b3bc（注意：测试为空壳，见二档）

## 二、部分实现（代码在，但有 bug 或测试缺失）

- [x] sorts/go 红测试修复（C1 完成）：bucket.go 桶序号误用 nBuckets 整除致全负输入越界（历史 14d87d4 曾修但未修对），且末尾回拷只扫 n 个桶（nBuckets 可大于 n）；counting.go 未按 min 偏移索引致负数越界（此前被 bucket panic 掩盖）。补 TestBucketSortSparse 稀疏两元素回归用例。
  验收达成：sorts 模块 `go test ./...` 7 个测试全绿。证据：本条修复 commit
- [x] go.work 收录全部 4 个 module（C2 完成）。顺带最小修复 linkedlist 空壳测试的编译错误（declared and not used），使四模块测试命令均可执行。
  验收达成：array/sorts/stack/linkedlist 四模块 `go test ./...` 均可执行。证据 3c39425
- [x] stack/go ArrayStack 补齐测试（C2 完成）。测试先红后修：Pop 原先不弹栈（等同 Top）、Flush 后复用 Size()=len(data) 失真、String 漏栈顶元素，三处全修，加编译期接口守卫 `var _ Stack`。
  验收达成：stack 模块 ArrayStack 3 测试全绿。证据 7e2f349
- [x] stack/go LinkedListStack 补 Push/Pop/Top/Size/Flush 五方法（C2 完成）。测试先行（编译红确认缺方法），头插法实现，空栈/ LIFO /Flush 复用三组用例 + 接口守卫。
  验收达成：stack 模块 6 测试全绿。证据 040b66c
- [x] linkedlist/go 测试全覆盖（C3 完成）。测试先行暴露 5 个真 bug 并全部修复：String 死循环（漏 cur=cur.next）、Contains 漏判尾节点、HasCycle 单节点 nil panic、FindMiddleNode 偶数长度 panic、Insert 持 RLock 写数据（-race 验证通过）。11 个测试 + 中点/判环边界用例。
  验收达成：linkedlist 模块 11 测试全绿（含 -race）。证据 f7a6b7a

## 三、只有设计 / 完全没有

- [x] docs/ 真实内容（C3 完成）：新增 docs/notes/ 四篇笔记（array/sorts/stack/linkedlist，含 API 表、复杂度、历史坑位回归记录）+ 索引页，接入 navbar/sidebar。vuepress beta.53 在 Node 24 下构建可用。
  验收达成：`npm run docs:build` 成功，13→18 页全渲染。证据 666a1ce
- [x] queue / tree / graph / heap / hash 结构（C4 完成）：全部按「module 骨架 + 基础实现 + 测试」落地，一个结构一个 commit，并收录进 go.work：
  - queue：Queue 接口 + 数组队列（头游标）/链表队列，6 测试 —— 证据 c231579
  - tree：泛型 BST（插入/查/删三态/最值/中序），6 测试 —— 证据 08c57bf
  - heap：泛型最小堆（上浮/下沉/O(n) 建堆），6 测试 —— 证据 a64bf41
  - hash：链地址哈希表（FNV-1a + 负载因子倍增扩容），5 测试 —— 证据 d7a45a6（模块名取 hashmap 避让标准库）
  - graph：邻接表有向/无向图 + BFS/DFS（判重边/环/连通分量用例），6 测试 —— 证据 1d1007d
  验收达成：全仓 9 模块 `go test ./...` 全绿。

## 四、后续可选（新账，非本轮验收）

- [ ] docs/notes/ 为 queue/tree/heap/hash/graph 五个新模块补笔记
- [ ] BST 增加高度/平衡检查；graph 增加 BFS 最短路（边权 1）
