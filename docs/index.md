---
layout: home

hero:
  name: Algorithm Notes
  text: Data structures and algorithms in Go
  tagline: Nine small Go modules with unit tests, covering arrays, sorting, stacks, linked lists, queues, binary search trees, heaps, hash maps, and graphs.
  image:
    src: /logo.svg
    alt: Algorithm Notes
  actions:
    - theme: brand
      text: Browse notes
      link: /notes/
    - theme: alt
      text: 简体中文
      link: /zh/

features:
  - title: Tested implementations
    details: Every data structure and sorting implementation ships with unit tests; run a module with `cd <module>/go && go test ./...`.
  - title: Generics where useful
    details: Tree and heap are generic; the array wrapper is generic too, with explicit bounds checks.
  - title: Consistent structure
    details: Each note documents the source location, public interface, usage example, and test status.
  - title: Bilingual
    details: English lives at the site root; Simplified Chinese lives under `/zh/`.
---

## Notes

| Topic | Module | Implementation |
| --- | --- | --- |
| [Array](/notes/array) | `array/go` | generic array + bounds check |
| [Sorting](/notes/sorts) | `sorts/go` | bubble / insertion / selection / merge / quick / counting / bucket |
| [Stack](/notes/stack) | `stack/go` | array stack / linked-list stack |
| [Linked List](/notes/linkedlist) | `linkedlist/go` | head sentinel + RWLock + fast/slow pointers |
| [Queue](/notes/queue) | `queue/go` | array queue (head cursor) / linked-list queue |
| [Binary Search Tree](/notes/tree) | `tree/go` | generic BST: insert / search / remove / min-max / in-order / height / balance |
| [Binary Heap](/notes/heap) | `heap/go` | generic min-heap: sift-up / sift-down / O(n) heapify |
| [Hash Map](/notes/hash) | `hash/go` | chaining + FNV-1a + load-factor doubling resize |
| [Graph](/notes/graph) | `graph/go` | adjacency-list directed/undirected graph + BFS / DFS / shortest path |

