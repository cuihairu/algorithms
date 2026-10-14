---
title: 栈
icon: layer-group
---

# 栈

> [English](stack.md) | **中文**

`stack/go` 模块定义统一 `Stack` 接口，提供数组与链表两种实现。

## 实现位置

- 源码：`stack/go/stack.go`（接口）、`stack/go/array.go`、`stack/go/linkedlist.go`
- 测试：`stack/go/array_test.go`、`stack/go/linkedlist_test.go`

## 接口

```go
type Stack interface {
    Push(v any)
    Pop() any
    IsEmpty() bool
    Size() int
    Top() any
    Flush()
}
```

## 实现

| 实现 | 结构 | 特点 |
| --- | --- | --- |
| `ArrayStack` | 切片 + top 游标 | `Flush` 后复用底层数组；`String()` 输出 `[a<b<c]`（底→顶） |
| `LinkedListStack` | 单链表头插 | 无扩容；节点按需分配 |

两者空栈 `Pop`/`Top` 均返回 `nil`，不 panic。

## 历史坑位（已修复，回归用例在测试里）

- **ArrayStack**：`Pop` 原先不移动 top（等同 `Top`）；`Size` 用 `len(data)` 导致 `Flush` 后复用时逻辑大小失真；`String` 漏打栈顶元素。
- 两个实现均加了编译期接口守卫：`var _ Stack = (*ArrayStack)(nil)`。

## 测试状态

```bash
cd stack/go && go test ./...
```

当前 6 个测试全部通过（空栈 / LIFO / Flush 复用 × 两种实现）。
