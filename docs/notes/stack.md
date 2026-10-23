---
title: Stack
icon: layer-group
---

# Stack

> **English** | [中文](/zh/notes/stack)

`stack/go` defines a unified `Stack` interface and provides array and linked list implementations.

## Location

- Source: `stack/go/stack.go` (interface), `stack/go/array.go`, `stack/go/linkedlist.go`
- Tests: `stack/go/array_test.go`, `stack/go/linkedlist_test.go`

## Interface

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

## Implementations

| Implementation | Structure | Notes |
| --- | --- | --- |
| `ArrayStack` | slice + top cursor | `Flush` reuses the underlying array; `String()` outputs `[a<b<c]` (bottom → top) |
| `LinkedListStack` | singly linked list, head insert | no resizing; nodes allocated on demand |

Both implementations return `nil` for `Pop`/`Top` on an empty stack, no panic.

## Historical pitfalls (fixed, regression cases live in tests)

- **ArrayStack**: `Pop` originally did not move `top` (equivalent to `Top`); `Size` used `len(data)`, which distorted the logical size when reusing after `Flush`; `String` omitted the top element.
- Both implementations have compile-time interface guards: `var _ Stack = (*ArrayStack)(nil)`.

## Test status

```bash
cd stack/go && go test ./...
```

All 6 tests pass (empty stack / LIFO / Flush reuse × two implementations).
