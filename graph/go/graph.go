package graph

// Graph 邻接表无权图,顶点为字符串,支持有向/无向
type Graph struct {
	adj      map[string][]string
	directed bool
	edges    int
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		adj:      make(map[string][]string),
		directed: directed,
	}
}

// AddVertex 添加顶点,已存在则不动作
func (this *Graph) AddVertex(v string) {
	if _, ok := this.adj[v]; !ok {
		this.adj[v] = nil
	}
}

// AddEdge 添加边,端点不存在时自动补顶点,重复边不动作
func (this *Graph) AddEdge(from, to string) {
	if this.hasEdge(from, to) {
		return
	}
	this.AddVertex(from)
	this.AddVertex(to)
	this.adj[from] = append(this.adj[from], to)
	this.edges++
	if !this.directed {
		this.adj[to] = append(this.adj[to], from)
	}
}

func (this *Graph) hasEdge(from, to string) bool {
	for _, v := range this.adj[from] {
		if v == to {
			return true
		}
	}
	return false
}

func (this *Graph) VertexCount() int {
	return len(this.adj)
}

func (this *Graph) EdgeCount() int {
	return this.edges
}

func (this *Graph) Neighbors(v string) []string {
	out := make([]string, len(this.adj[v]))
	copy(out, this.adj[v])
	return out
}

// HasEdge 判断边是否存在
func (this *Graph) HasEdge(from, to string) bool {
	return this.hasEdge(from, to)
}

// BFS 广度优先遍历,返回从 start 可达的访问序
func (this *Graph) BFS(start string) []string {
	if _, ok := this.adj[start]; !ok {
		return nil
	}
	visited := map[string]bool{start: true}
	queue := []string{start}
	out := make([]string, 0, len(this.adj))
	for head := 0; head < len(queue); head++ {
		cur := queue[head]
		out = append(out, cur)
		for _, next := range this.adj[cur] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	return out
}

// DFS 深度优先遍历(先序),返回从 start 可达的访问序
func (this *Graph) DFS(start string) []string {
	if _, ok := this.adj[start]; !ok {
		return nil
	}
	visited := make(map[string]bool)
	out := make([]string, 0, len(this.adj))
	var visit func(v string)
	visit = func(v string) {
		visited[v] = true
		out = append(out, v)
		for _, next := range this.adj[v] {
			if !visited[next] {
				visit(next)
			}
		}
	}
	visit(start)
	return out
}

// ShortestPath 返回 from 到 to 的最短路径(边权 1),含两端点;起点即终点返回单点,不可达或顶点不存在返回 nil
func (this *Graph) ShortestPath(from, to string) []string {
	if _, ok := this.adj[from]; !ok {
		return nil
	}
	if _, ok := this.adj[to]; !ok {
		return nil
	}
	if from == to {
		return []string{from}
	}
	// BFS 记父节点,首次命中 to 时回溯路径
	parent := map[string]string{from: ""}
	queue := []string{from}
	for head := 0; head < len(queue); head++ {
		cur := queue[head]
		for _, next := range this.adj[cur] {
			if _, seen := parent[next]; seen {
				continue
			}
			parent[next] = cur
			if next == to {
				path := []string{to}
				for p := cur; p != ""; p = parent[p] {
					path = append(path, p)
				}
				// 当前为 to...from,原地反转
				for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
					path[i], path[j] = path[j], path[i]
				}
				return path
			}
			queue = append(queue, next)
		}
	}
	return nil
}
