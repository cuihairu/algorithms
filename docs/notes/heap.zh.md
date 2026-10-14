---
title: 二叉堆
icon: sort-amount-up
---

# 二叉堆

> [English](heap.md) | **中文**

`heap/go` 模块实现泛型最小堆，数组存储，父节点 ≤ 两个子节点。

## 实现位置

- 源码：`heap/go/heap.go`
- 测试：`heap/go/heap_test.go`

## API

```go
func NewHeap[T constraints.Ordered]() *Heap[T]
func NewHeapFrom[T constraints.Ordered](vals []T) *Heap[T] // O(n) 建堆
func (this *Heap[T]) Size() int
func (this *Heap[T]) IsEmpty() bool
func (this *Heap[T]) Peek() (T, bool)  // 空堆返回 false
func (this *Heap[T]) Push(v T)
func (this *Heap[T]) Pop() (T, bool)   // 空堆返回 false
```

## 实现要点

- 索引公式：父 `(i-1)/2`，子 `2i+1` / `2i+2`。
- `Push` 追加到末尾后**上浮**（siftUp）；`Pop` 用末元素覆盖堆顶后**下沉**（siftDown），均只与一侧子节点比较换更小者。
- `NewHeapFrom` 自底向上堆化：从最后一个非叶节点 `len/2-1` 逆序 `siftDown`，整体 O(n)，优于逐个 Push 的 O(n log n)。建堆时 `copy` 输入切片，不改调用方数据。
- 出参统一 `(value, ok)`，空堆返回零值 + `false`，不 panic。

## 复杂度

| 操作 | 复杂度 |
| --- | --- |
| Peek | O(1) |
| Push / Pop | O(log n) |
| NewHeapFrom | O(n) |

## 测试状态

```bash
cd heap/go && go test ./...
```

当前 6 个测试全部通过（空堆 / Push-Pop 有序 / 交错操作 / 切片建堆 / 空切片建堆 / 字符串泛型）。
