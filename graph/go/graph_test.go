package graph

import (
	"reflect"
	"testing"
)

func eqSeq(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

func TestGraphEmpty(t *testing.T) {
	g := NewGraph(true)
	if g.VertexCount() != 0 {
		t.Fatalf("vertex count = %d, want 0", g.VertexCount())
	}
	if g.BFS("a") != nil {
		t.Fatal("BFS on absent start should be nil")
	}
	if g.DFS("a") != nil {
		t.Fatal("DFS on absent start should be nil")
	}
}

func TestGraphAddVertexEdge(t *testing.T) {
	g := NewGraph(true)
	g.AddVertex("a")
	g.AddVertex("a") // 重复添加不动作
	if g.VertexCount() != 1 {
		t.Fatalf("vertex count = %d, want 1", g.VertexCount())
	}

	g.AddEdge("a", "b")
	if g.VertexCount() != 2 {
		t.Fatalf("vertex count = %d, want 2", g.VertexCount())
	}
	if g.EdgeCount() != 1 {
		t.Fatalf("edge count = %d, want 1", g.EdgeCount())
	}
	if !g.HasEdge("a", "b") {
		t.Fatal("edge a->b should exist")
	}
	if g.HasEdge("b", "a") {
		t.Fatal("directed graph should not have reverse edge")
	}

	g.AddEdge("a", "b") // 重复边不动作
	if g.EdgeCount() != 1 {
		t.Fatalf("edge count = %d, want 1", g.EdgeCount())
	}

	// 有向图:b 不可达 a
	if got := g.BFS("b"); !eqSeq(got, []string{"b"}) {
		t.Fatalf("BFS from b = %v, want [b]", got)
	}
}

func TestGraphBFSDirected(t *testing.T) {
	// a -> b,c; b -> d; c -> e
	g := NewGraph(true)
	g.AddEdge("a", "b")
	g.AddEdge("a", "c")
	g.AddEdge("b", "d")
	g.AddEdge("c", "e")

	if got := g.BFS("a"); !eqSeq(got, []string{"a", "b", "c", "d", "e"}) {
		t.Fatalf("BFS = %v, want [a b c d e]", got)
	}
	if got := g.DFS("a"); !eqSeq(got, []string{"a", "b", "d", "c", "e"}) {
		t.Fatalf("DFS = %v, want [a b d c e]", got)
	}
	// 不可达分支
	if got := g.BFS("d"); !eqSeq(got, []string{"d"}) {
		t.Fatalf("BFS from d = %v, want [d]", got)
	}
}

func TestGraphUndirected(t *testing.T) {
	g := NewGraph(false)
	g.AddEdge("a", "b")
	g.AddEdge("a", "c")

	if !g.HasEdge("b", "a") || !g.HasEdge("c", "a") {
		t.Fatal("undirected edge should exist both ways")
	}
	if g.EdgeCount() != 2 {
		t.Fatalf("edge count = %d, want 2 (logical)", g.EdgeCount())
	}
	if got := g.BFS("b"); !eqSeq(got, []string{"b", "a", "c"}) {
		t.Fatalf("BFS from b = %v, want [b a c]", got)
	}
	if n := g.Neighbors("a"); !eqSeq(n, []string{"b", "c"}) {
		t.Fatalf("neighbors of a = %v, want [b c]", n)
	}
}

// 回归:无向 AddEdge(A,A) 曾向自身邻接表追加两次,Neighbors(a)=[a a];自环应只记一次。
func TestGraphUndirectedSelfLoop(t *testing.T) {
	g := NewGraph(false)
	g.AddEdge("a", "a")
	if g.EdgeCount() != 1 {
		t.Fatalf("edge count = %d, want 1", g.EdgeCount())
	}
	if n := g.Neighbors("a"); !eqSeq(n, []string{"a"}) {
		t.Fatalf("neighbors of a = %v, want [a]", n)
	}
	if !g.HasEdge("a", "a") {
		t.Fatal("self-loop edge should exist")
	}
}

func TestGraphDisconnected(t *testing.T) {
	g := NewGraph(true)
	g.AddEdge("a", "b")
	g.AddEdge("x", "y")

	// BFS 不跨连通分量
	if got := g.BFS("a"); !eqSeq(got, []string{"a", "b"}) {
		t.Fatalf("BFS from a = %v, want [a b]", got)
	}
	if got := g.DFS("x"); !eqSeq(got, []string{"x", "y"}) {
		t.Fatalf("DFS from x = %v, want [x y]", got)
	}
}

func TestGraphShortestPath(t *testing.T) {
	// a -> b -> d -> f;a -> c -> e
	g := NewGraph(true)
	g.AddEdge("a", "b")
	g.AddEdge("b", "d")
	g.AddEdge("d", "f")
	g.AddEdge("a", "c")
	g.AddEdge("c", "e")

	// 最短路径取更少边数的路线
	if got := g.ShortestPath("a", "f"); !eqSeq(got, []string{"a", "b", "d", "f"}) {
		t.Fatalf("shortest a->f = %v, want [a b d f]", got)
	}
	if got := g.ShortestPath("a", "e"); !eqSeq(got, []string{"a", "c", "e"}) {
		t.Fatalf("shortest a->e = %v, want [a c e]", got)
	}
	// 起点即终点
	if got := g.ShortestPath("a", "a"); !eqSeq(got, []string{"a"}) {
		t.Fatalf("shortest a->a = %v, want [a]", got)
	}
	// 有向图反向不可达
	if got := g.ShortestPath("f", "a"); got != nil {
		t.Fatalf("shortest f->a = %v, want nil", got)
	}
	// 顶点不存在
	if got := g.ShortestPath("a", "z"); got != nil {
		t.Fatalf("shortest a->z = %v, want nil", got)
	}
	if got := g.ShortestPath("z", "a"); got != nil {
		t.Fatalf("shortest z->a = %v, want nil", got)
	}
}

func TestGraphShortestPathDirectBeatsIndirect(t *testing.T) {
	// 直达边必须胜出:a->c 直达 vs a->b->c
	g := NewGraph(true)
	g.AddEdge("a", "b")
	g.AddEdge("b", "c")
	g.AddEdge("a", "c")

	if got := g.ShortestPath("a", "c"); !eqSeq(got, []string{"a", "c"}) {
		t.Fatalf("shortest a->c = %v, want [a c]", got)
	}
}

func TestGraphShortestPathUndirected(t *testing.T) {
	g := NewGraph(false)
	g.AddEdge("a", "b")
	g.AddEdge("b", "c")

	if got := g.ShortestPath("a", "c"); !eqSeq(got, []string{"a", "b", "c"}) {
		t.Fatalf("shortest a->c = %v, want [a b c]", got)
	}
	if got := g.ShortestPath("c", "a"); !eqSeq(got, []string{"c", "b", "a"}) {
		t.Fatalf("shortest c->a = %v, want [c b a]", got)
	}
	// 不连通分量不可达
	g.AddEdge("x", "y")
	if got := g.ShortestPath("a", "x"); got != nil {
		t.Fatalf("shortest a->x = %v, want nil", got)
	}
}

func TestGraphCycle(t *testing.T) {
	// 有环图遍历必须终止且不重复访问
	g := NewGraph(true)
	g.AddEdge("a", "b")
	g.AddEdge("b", "c")
	g.AddEdge("c", "a")

	if got := g.BFS("a"); !eqSeq(got, []string{"a", "b", "c"}) {
		t.Fatalf("BFS on cycle = %v, want [a b c]", got)
	}
	if got := g.DFS("a"); !eqSeq(got, []string{"a", "b", "c"}) {
		t.Fatalf("DFS on cycle = %v, want [a b c]", got)
	}
}
