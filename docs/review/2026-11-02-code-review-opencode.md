---
title: 代码深审报告（2026-11-02，opencode）
---

# 代码深审报告

> 依据 requesting-code-review 流程（安全静态扫描 → 隐藏 bug 猎捕 → 质量门 → 分级修复），本仓全量。日期 2026-11-02。
> 扫描范围：`array` / `sorts` / `stack` / `linkedlist` / `queue` / `tree` / `heap` / `hash` / `graph` 共 9 个 Go module 全部源码，以及 docs 站构建。

## ① 安全静态扫描（全仓）

| 检查项 | 结果 |
| --- | --- |
| 硬编码密钥（api_key/secret/password/token） | 无 |
| Shell 注入 / `exec.Command`+shell | 无 |
| SQL 拼接 | 无（纯数据结构库，无 DB/网络/文件 IO 面） |
| `eval` / `exec` | 无 |
| 不安全反序列化（`pickle` 等） | 无 |
| 路径穿越 / 文件 IO（`..` / `os.Open` / `filepath.Join`） | 无 |
| `gosec` | 1 条告警（见 B9，整数转换，误报） |
| `govulncheck` | 代码路径 0 漏洞（依赖层 13-14 条未调用） |

结论：安全面干净；唯一 gosec 告警为哈希表桶数掩码转换，桶数恒为 2 的幂且远小于 `uint32` 上限，无实际溢出。

## ② 隐藏 bug 猎捕（按严重度）

| ID | 严重度 | 位置 | 缺陷 | 证据 | 状态 |
| --- | --- | --- | --- | --- | --- |
| B1 | P0 | sorts/go/merge.go:12-19 | `merge` 取左侧元素分支缺 `i++`：左 ≤ 右时死循环越界 | `MergeSort([1,2,3,4,5])` panic `index out of range [2] with length 2`（历史实证） | **已修** `9edeae7` |
| B2 | P0 | sorts/go/merge.go:20-23 | 缺右半段余量回拷循环：左半先耗尽时尾部零值写回 | `MergeSort([1,2,3,4,5])` → `[1,0,0,0,0]`（历史实证） | **已修** `9edeae7` |
| B3 | P1 | array/go/array.go:36-40 | `isIndexOutOfRange` 用 `Cap()` 而非 `Len()`：`[len,cap)` 被当作合法 | `Find(5)` 零值无错；`Insert(6,99)` 产出 `[1,2,0]` 空洞；`Pop(5)` 无中生有减 Len（历史实证） | **已修** `0c4cabd` |
| B4 | P2 | sorts/*_test.go | 冒泡/计数/插入/快排测试仅 `t.Log` 无断言（`selection_test.go` 为空文件），烟测掩盖了 B1/B2 | 仅 `bucket_test.go`/`merge_test.go` 有 `IntsAreSorted`/`DeepEqual` 断言 | 记 todo |
| B5 | P2 | graph/go/graph.go:24-37 | 无向 `AddEdge(A,A)` 自环向 `adj[A]` 追加两次（`EdgeCount` 记 1） | `AddEdge("a","a")` → `Neighbors("a")=[a a]`，`EdgeCount()=1`（实证） | 记 todo |
| B6 | P3 | linkedlist/go/linkedlist.go:105-108 | `RemoveLast` 重复 `isEmpty` 判断（死代码） | 连续两次 `if this.isEmpty()` | 记 todo |
| B7 | P3 | queue/go/array.go、queue/go/linkedlist.go | `Queue` 接口无源码级编译期守卫 `var _ Queue`（仅测试文件有） | grep 源码无 `var _ Queue` | 记 todo |
| B8 | P3 | queue/stack/tree/heap/hash/graph | 仅 linkedlist 文档化 RWLock；其余未声明非并发安全 | 各结构无锁 | 记 todo |
| B9 | P3 | hash/go/hashmap.go:103-105 | gosec G115：`int(fnv1a(key) & uint32(len(this.buckets)-1))` 整数转换告警 | gosec 实证 1 条（HIGH/MEDIUM）；桶数为 2 的幂，无实际溢出，误报 | 记 todo |
| B10 | P3 | array/go/array.go:32-43 | `Insert` 先判 `FullArray` 再判 `OutOfRange`：满数组越界插入返回 `FullArray` 而非 `OutOfRange` | `NewArray[int](2)` 填满后 `Insert(5,9)` 返回 `full array`（实证） | 记 todo |
| B11 | P3 | linkedlist/go/linkedlist.go:172-186 | `Insert` 对 `index > size+1` 静默追加到尾部，无越界反馈 | `Insert(10,99)` 对 5 元素链表追加成功（实证） | 记 todo |
| B12 | P3 | queue/go/array.go:27-30 | `ArrayQueue.Enqueue` 始终 `append`，非空稳态下 `len(data)` 随总入队次数增长，游标前部槽位不复用 | 长期稳态负载下内存 O(总操作数) 而非 O(当前大小) | 记 todo |

**猎捕结论**：本轮未发现新的 P0/P1；2 个 P0、1 个 P1 均已在历史 commit 中修复；其余 P2×2、P3×7 记 todo。对 `sorts` 7 种排序与 `graph.ShortestPath` 做了随机/边界/暴力对照 property 测试，未发现新缺陷。

## ③ 质量门

| 门 | 结果 |
| --- | --- |
| `go test ./...`（9 模块） | 全绿 |
| `linkedlist -race` | 绿 |
| 全模块 `-race` | 绿 |
| `go vet ./...`（9 模块） | 干净 |
| `staticcheck` | 仅风格：ST1006（`this` 接收者）、ST1012（错误变量命名）、S1019（make 长度）、S1008；为代码库一致风格，非缺陷 |
| `gosec` | 1 条（B9，误报） |
| `govulncheck` | 代码路径 0 漏洞 |
| `gofmt -l` | 无差异 |
| `pnpm run docs:build` | 成功（37.6s） |
| 排序正确性探针（property，7 种排序 × 2000 轮随机 + 边界） | 全部通过 |
| 图最短路暴力对照（500 轮随机图） | 全部通过 |

## ④ 修复与剩余

- **发现**：P0×2（B1/B2）、P1×1（B3）、P2×2（B4/B5）、P3×7（B6-B12）
- **已修**：3（B1/B2 合并为一 commit `9edeae7`；B3 为 `0c4cabd`），每项测试绿
- **本轮新增 P0/P1 修复**：0（未发现新的 P0/P1）
- **剩余（记 todo）**：P2×2、P3×7

## ⑤ 总结

- 发现数：12
- 已修数：3
- 剩余数：9
- commit 清单：
  - `9edeae7` fix: merge sort infinite loop and missing right-half copy
  - `0c4cabd` fix: array bounds checks use logical length not capacity
  - (本报告) docs: add code review report 2026-11-02 (opencode)
- push：本轮报告提交后推送至 origin/main
