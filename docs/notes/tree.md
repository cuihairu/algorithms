---
title: Binary Search Tree
icon: tree
---

# Binary Search Tree

> **English** | [中文](tree.zh.md)

`tree/go` implements a generic binary search tree (BST) with no duplicate elements.

## Location

- Source: `tree/go/bst.go`
- Tests: `tree/go/bst_test.go`

## API

```go
func NewBST[T constraints.Ordered]() *BST[T]
func (this *BST[T]) Size() int
func (this *BST[T]) IsEmpty() bool
func (this *BST[T]) Insert(v T) bool   // duplicate is a no-op and returns false
func (this *BST[T]) Search(v T) bool
func (this *BST[T]) Remove(v T) bool   // returns false if absent
func (this *BST[T]) Min() (T, bool)    // empty tree returns false
func (this *BST[T]) Max() (T, bool)
func (this *BST[T]) InOrder() []T      // ascending slice
func (this *BST[T]) Height() int       // height in nodes, empty tree is 0
func (this *BST[T]) IsBalanced() bool  // left/right subtree height diff ≤ 1 at every node
```

## Implementation notes

- Insert/remove/search are all recursive and return "new subtree root + whether it took effect", with `size` maintained uniformly by the caller.
- Removal splits into three cases: leaf/single child are detached directly; **a node with two children is replaced by the smallest node of its right subtree (the successor)**, which is then removed from the right subtree.
- `InOrder` walks recursively in a closure with capacity preallocated by `size`.
- `Height` and `IsBalanced` share one pass: `checkBalance` returns both the subtree height and the balance at once, O(n) in one go, avoiding repeated recomputation of `Height`.

## Complexity

| Operation | Average | Worst (degenerates to a list) |
| --- | --- | --- |
| Insert / Search / Remove | O(log n) | O(n) |
| Min / Max | O(h) | O(n) |
| InOrder | O(n) | O(n) |
| Height / IsBalanced | O(n) | O(n) |

## Test status

```bash
cd tree/go && go test ./...
```

All 8 tests pass (empty tree / insert dedup / Min/Max/in-order / three-case removal / root removal variants / string generics / height & balance).
