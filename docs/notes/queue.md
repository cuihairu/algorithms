---
title: Queue
icon: stream
---

# Queue

> **English** | [中文](/zh/notes/queue)

`queue/go` defines a unified `Queue` interface and provides array and linked list implementations, first in first out (FIFO).

## Location

- Source: `queue/go/queue.go` (interface), `queue/go/array.go`, `queue/go/linkedlist.go`
- Tests: `queue/go/array_test.go`, `queue/go/linkedlist_test.go`

## Interface

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

## Implementations

Constructors: `NewArrayQueue()` and `NewLinkedListQueue()`.

| Implementation | Structure | Notes |
| --- | --- | --- |
| `ArrayQueue` | slice + `head` front cursor | `Dequeue` only moves the cursor, never clears elements; the underlying slice is compacted forward when at least half is consumed, and is reclaimed plus the cursor reset **when the queue is emptied**; `String()` outputs `[a<b<c]` (front → tail) |
| `LinkedListQueue` | singly linked list with head/tail pointers | enqueue via tail, dequeue via head, tail maintained automatically |

Both implementations return `nil` for `Dequeue`/`Front` on an empty queue, no panic. While the `ArrayQueue` is not empty, `Enqueue` compacts the backing slice forward once at least half has been consumed, keeping memory bounded.

## Complexity

| Operation | ArrayQueue | LinkedListQueue |
| --- | --- | --- |
| Enqueue / Dequeue / Front | O(1) | O(1) |

## Test status

```bash
cd queue/go && go test ./...
```

All 6 tests pass (empty queue / FIFO order / Flush reuse × two implementations).
