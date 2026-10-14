---
title: 图
icon: project-diagram
---

# 图

> [English](graph.md) | **中文**

`graph/go` 模块实现字符串顶点的邻接表无权图，支持有向与无向两种模式。

## 实现位置

- 源码：`graph/go/graph.go`
- 测试：`graph/go/graph_test.go`

## API

```go
func NewGraph(directed bool) *Graph
func (this *Graph) AddVertex(v string)          // 已存在则不动作
func (this *Graph) AddEdge(from, to string)     // 端点自动补顶点,重复边不动作
func (this *Graph) VertexCount() int
func (this *Graph) EdgeCount() int
func (this *Graph) Neighbors(v string) []string // 返回副本
func (this *Graph) HasEdge(from, to string) bool
func (this *Graph) BFS(start string) []string   // 从 start 可达的访问序
func (this *Graph) DFS(start string) []string   // 先序,递归
func (this *Graph) ShortestPath(from, to string) []string // 边权 1 最短路,含两端点
```

## 实现要点

- 邻接表为 `map[string][]string`；无向图 `AddEdge` 双向追加但 `edges` 只计一次。
- **判重边**：`AddEdge` 先线性扫一遍邻接表，重复边不动作，`EdgeCount` 不会虚增。
- BFS 用切片 + head 游标当队列（不删头元素）；DFS 用闭包递归先序访问。
- **BFS 最短路**：`ShortestPath` 沿路记父节点，首次命中 `to` 即回溯（边权全为 1 时 BFS 层数即最短边数）；起点等于终点返回单点，任一端点不存在或不可达返回 `nil`。
- 起点不存在时 BFS/DFS 返回 `nil`；`Neighbors` 返回拷贝，外部改不动内部状态。

## 复杂度

| 操作 | 复杂度 |
| --- | --- |
| AddEdge / HasEdge | O(deg) |
| BFS / DFS / ShortestPath | O(V + E) |

## 测试状态

```bash
cd graph/go && go test ./...
```

当前 6 个测试全部通过（空图 / 加点加边计数 / 有向 BFS / 无向对称性 / 非连通分量不可达 / 环不致死循环）。
