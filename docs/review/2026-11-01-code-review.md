---
title: 代码深审报告（2026-11-01）
---

# 代码深审报告

> 依据 requesting-code-review 技能流程（安全静态扫描 → 隐藏 bug 猎捕 → 质量门 → 分级修复），本仓全量。日期 2026-11-01。
> 扫描范围：`array` / `sorts` / `stack` / `linkedlist` / `queue` / `tree` / `heap` / `hash` / `graph` 共 9 个 Go module 全部源码。

## ① 安全静态扫描（全仓）

- 硬编码密钥（api_key/secret/password/token）：**无**
- Shell 注入 / `os.system` / `exec.Command`+shell：**无**
- SQL 拼接（`execute(f"` / `.format(...SELECT`)：**无**
- `eval` / `exec`：**无**
- 不安全反序列化（`pickle.loads` 等）：**无**
- 路径穿越 / 文件 IO（`..` / `filepath.Join` / `os.Open`）：**无**（纯数据结构库，无 I/O/网络/SQL 面）
- `gosec`（sorts/array 全量）：**0 告警**

结论：安全面干净，无安全缺陷。

## ② 隐藏 bug 猎捕（按严重度）

| ID | 严重度 | 位置 | 缺陷 | 证据 | 状态 |
| --- | --- | --- | --- | --- | --- |
| B1 | P0 | sorts/go/merge.go:12-19 | `merge` 取左侧元素分支缺 `i++`：任何「左元素 ≤ 右元素」输入死循环越界 | `MergeSort([1,2,3,4,5])` panic `index out of range [2] with length 2`（实证） | **已修** `9edeae7` |
| B2 | P0 | sorts/go/merge.go:20-23 | 缺右半段余量回拷循环：左半先耗尽时尾部零值被写回 | `MergeSort([1,2,3,4,5])` → `[1,0,0,0,0]`（实证） | **已修** `9edeae7` |
| B3 | P1 | array/go/array.go:36-40 | `isIndexOutOfRange` 用 `Cap()` 而非 `Len()`：`[len,cap)` 区间被当作合法 | `Find(5)` 零值无错；`Insert(6,99)` 产出 `[1,2,0]` 空洞；`Pop(5)` 无中生有减 Len（实证） | **已修** `0c4cabd` |
| B4 | P2 | sorts/go/*_test.go | 冒泡/计数/插入/快排测试仅 `t.Log` 无断言（selection_test.go 空文件），烟测掩盖了 B1/B2 | 仅 `bucket_test.go` 有 `IntsAreSorted` 断言 | 记 todo |
| B5 | P2 | graph/go/graph.go:39-47 | 无向 `AddEdge(A,A)` 自环向 `adj[A]` 追加两次（`EdgeCount` 记 1） | `AddEdge` 无向分支无条件 append | 记 todo |
| B6 | P3 | linkedlist/go/linkedlist.go:105-108 | `RemoveLast` 重复 `isEmpty` 判断（死代码） | 连续两次 `if this.isEmpty()` | 记 todo |
| B7 | P3 | queue/go/array.go、queue/go/linkedlist.go | `Queue` 接口无编译期守卫 `var _ Queue`（stack 已有） | grep 无 `var _ Queue` | 记 todo |
| B8 | P3 | queue/stack/tree/heap/hash/graph | 仅 linkedlist 文档化 RWLock，其余未声明非并发安全 | 各结构无锁 | 记 todo |

**猎捕结论**：2 个 P0（MergeSort 双重缺陷，被非断言烟测掩盖）、1 个 P1（array 边界用 Cap）、P2×2、P3×3。B1/B2/B3 当场修复；其余记 todo。

## ③ 质量门

| 门 | 结果 |
| --- | --- |
| `go test ./...`（9 模块） | 全绿 |
| `linkedlist -race` | 绿 |
| `go vet ./...`（9 模块） | 干净 |
| `staticcheck` | 仅风格：ST1006（`this` 接收者）、ST1012（错误变量命名）、S1019（make 长度）；为代码库一致风格，非缺陷 |
| `gosec` | 0 告警 |
| 排序正确性探针（7 种排序 × 9 种输入：升序/降序/重复/全等/单元素/空/负数/稠密负/稀疏） | 全部通过（修复后） |

## ④ 修复与剩余

- **发现**：P0×2（B1/B2）、P1×1（B3）、P2×2、P3×3（另风格 P3×1 忽略）
- **已修**：3（B1/B2 合并为一 commit `9edeae7`；B3 为 `0c4cabd`），每项测试绿
- **剩余（记 todo）**：P2×2、P3×3
- **commit 清单**：`9edeae7` fix: merge sort infinite loop and missing right-half copy；`0c4cabd` fix: array bounds checks use logical length not capacity
