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
