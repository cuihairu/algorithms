---
title: 队列
icon: stream
---

# 队列

> [English](queue.md) | **中文**

`queue/go` 模块定义统一 `Queue` 接口，提供数组与链表两种实现，先进先出（FIFO）。

## 实现位置

- 源码：`queue/go/queue.go`（接口）、`queue/go/array.go`、`queue/go/linkedlist.go`
- 测试：`queue/go/array_test.go`、`queue/go/linkedlist_test.go`

## 接口

```go
type Queue interface {
    Enqueue(v any)
    Dequeue() any
    IsEmpty() bool
    Size() int
    Front() any
    Flush()
}
```

## 实现

| 实现 | 结构 | 特点 |
| --- | --- | --- |
| `ArrayQueue` | 切片 + `head` 队首游标 | `Dequeue` 只右移游标不清元素；队列**清空时**才回收底层数组并归零游标；`String()` 输出 `[a<b<c]`（头→尾） |
| `LinkedListQueue` | 单链表 head/tail 双指针 | 入队走 tail、出队走 head，尾节点自动维护 |

两种实现空队 `Dequeue`/`Front` 均返回 `nil`，不 panic。注意 `ArrayQueue` 在队列始终非空时，底层切片会随 `Enqueue` 持续增长（只在清空时回收）。

## 复杂度

| 操作 | ArrayQueue | LinkedListQueue |
| --- | --- | --- |
| Enqueue / Dequeue / Front | O(1) | O(1) |

## 测试状态

```bash
cd queue/go && go test ./...
```

当前 6 个测试全部通过（空队 / FIFO 序 / Flush 复用 × 两种实现）。
