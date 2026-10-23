---
title: Graph
icon: project-diagram
---

# Graph

> **English** | [中文](/zh/notes/graph)

`graph/go` implements an adjacency-list unweighted graph with string vertices, supporting both directed and undirected modes.

## Location

- Source: `graph/go/graph.go`
- Tests: `graph/go/graph_test.go`

## API

```go
func NewGraph(directed bool) *Graph
func (this *Graph) AddVertex(v string)          // no-op if it already exists
func (this *Graph) AddEdge(from, to string)     // adds missing endpoints automatically, duplicate edges are no-ops
func (this *Graph) VertexCount() int
func (this *Graph) EdgeCount() int
func (this *Graph) Neighbors(v string) []string // returns a copy
func (this *Graph) HasEdge(from, to string) bool
func (this *Graph) BFS(start string) []string   // visit order reachable from start
func (this *Graph) DFS(start string) []string   // pre-order, recursive
func (this *Graph) ShortestPath(from, to string) []string // shortest path, edge weight 1, includes both endpoints
```

## Implementation notes

- The adjacency list is a `map[string][]string`; an undirected `AddEdge` appends in both directions but counts `edges` once.
- **Duplicate edge check**: `AddEdge` linearly scans the adjacency list first, so duplicate edges are no-ops and `EdgeCount` is not inflated.
- BFS uses a slice + head cursor as the queue (elements are never removed from the head); DFS recurses in a closure, pre-order.
- **BFS shortest path**: `ShortestPath` records parents along the way and backtracks on the first hit of `to` (with all edge weights equal to 1, the BFS level is the shortest edge count); start equals end returns a single point; if either endpoint is missing or unreachable, it returns `nil`.
- When the start is absent, BFS/DFS return `nil`; `Neighbors` returns a copy, so external changes do not affect internal state.

## Complexity

| Operation | Complexity |
| --- | --- |
| AddEdge / HasEdge | O(deg) |
| BFS / DFS / ShortestPath | O(V + E) |

## Test status

```bash
cd graph/go && go test ./...
```

All 9 tests pass (empty graph / vertex & edge counts / directed BFS / undirected symmetry / disconnected components unreachable / cycle terminates / shortest path / direct edge wins / undirected shortest path).
