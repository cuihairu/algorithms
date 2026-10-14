---
title: Binary Heap
icon: sort-amount-up
---

# Binary Heap

> **English** | [中文](heap.zh.md)

`heap/go` implements a generic min-heap backed by an array; every parent node is ≤ its two children.

## Location

- Source: `heap/go/heap.go`
- Tests: `heap/go/heap_test.go`

## API

```go
func NewHeap[T constraints.Ordered]() *Heap[T]
func NewHeapFrom[T constraints.Ordered](vals []T) *Heap[T] // O(n) heapify
func (this *Heap[T]) Size() int
func (this *Heap[T]) IsEmpty() bool
func (this *Heap[T]) Peek() (T, bool)  // empty heap returns false
func (this *Heap[T]) Push(v T)
func (this *Heap[T]) Pop() (T, bool)   // empty heap returns false
```

## Implementation notes

- Index formulas: parent `(i-1)/2`, children `2i+1` / `2i+2`.
- `Push` appends to the end and **sifts up**; `Pop` overwrites the top with the last element and **sifts down**, swapping with the smaller child.
- `NewHeapFrom` builds bottom-up: it sifts down from the last non-leaf node `len/2-1` in reverse order, O(n) overall, better than pushing one by one (O(n log n)). It copies the input slice, so the caller's data is untouched.
- Returns use `(value, ok)`; an empty heap returns the zero value + `false`, no panic.

## Complexity

| Operation | Complexity |
| --- | --- |
| Peek | O(1) |
| Push / Pop | O(log n) |
| NewHeapFrom | O(n) |

## Test status

```bash
cd heap/go && go test ./...
```

All 6 tests pass (empty heap / push-pop order / interleaved / slice heapify / empty slice heapify / string generics).
