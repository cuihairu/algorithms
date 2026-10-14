---
title: Notes
icon: note
---

# Algorithm Notes

> **English** | [中文](README.zh.md)

The data structures and sorting implementations in this repo are all Go (generics for `tree`/`heap`), all with unit tests.

| Topic | Module | Implementation |
| --- | --- | --- |
| [Array](array.md) | `array/go` | generic array + bounds check |
| [Sorting](sorts.md) | `sorts/go` | bubble / insertion / selection / merge / quick / counting / bucket |
| [Stack](stack.md) | `stack/go` | array stack / linked list stack |
| [Linked List](linkedlist.md) | `linkedlist/go` | head sentinel + RWLock + fast/slow pointers |
| [Queue](queue.md) | `queue/go` | array queue (head cursor) / linked list queue |
| [Binary Search Tree](tree.md) | `tree/go` | generic BST: insert / search / remove three cases / min-max / in-order / height / balance |
| [Binary Heap](heap.md) | `heap/go` | generic min-heap: sift-up / sift-down / O(n) heapify |
| [Hash Map](hash.md) | `hash/go` | chaining + FNV-1a + load-factor doubling resize |
| [Graph](graph.md) | `graph/go` | adjacency-list directed/undirected graph + BFS / DFS / shortest path |

## Run all the tests at once

All nine modules are included in `go.work`; run inside each module directory:

```bash
cd <module>/go && go test ./...
```
