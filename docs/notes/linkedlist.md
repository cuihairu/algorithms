---
title: Linked List
icon: link
---

# Linked List

> **English** | [中文](linkedlist.zh.md)

`linkedlist/go` implements a singly linked list with a head sentinel node, guarded by a read-write lock.

## Location

- Source: `linkedlist/go/linkedlist.go`
- Tests: `linkedlist/go/linkedlist_test.go` (11 test functions)

## Interface

| Method | Description |
| --- | --- |
| `AddFirst / AddLast` | head insert / tail insert |
| `GetFirst / GetLast` | read head / read tail |
| `RemoveFirst / RemoveLast` | remove head / remove tail, empty list returns nil |
| `Insert(index, element)` | 1-based index insert; `index<=0` is a no-op; out of range behaves as a tail insert |
| `DeleteByIndex(index)` | delete by index, out of range is a no-op |
| `DeleteByBottomIndex(index)` | delete the Nth from the end (fast/slow pointers) |
| `Contains(element)` | whether it contains the element |
| `Reverse()` | in-place reversal |
| `HasCycle()` | fast/slow pointer cycle detection |
| `FindMiddleNode()` | fast/slow pointers for the middle; for even length returns the earlier middle |
| `Size / IsEmpty / String` | routine |

## Historical pitfalls (fixed, regression cases live in tests)

- **String**: traversal forgot `cur = cur.next`, infinite loop.
- **Contains**: the loop condition `next.next != nil` missed the tail node.
- **HasCycle**: the condition lacked `fast != nil`, so a single-node list panicked on `nil.next`.
- **FindMiddleNode**: `fast.next.next` was not nil-checked, panicking on even length.
- **Insert**: held `RLock` while writing the list (data race), now `Lock`; verified with `-race`.

## Test status

```bash
cd linkedlist/go && go test -race ./...
```

All 11 tests pass (including a hand-built cyclic list and even-length middle lookup).
