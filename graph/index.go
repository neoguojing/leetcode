package graph

// Graph 邻接表
// 适合稀疏图
type Graph struct {
	Enum int
	Vnum int
	Adj  []Vertex
}

// Vertex 节点
type Vertex struct {
	Val       int
	FirstEdge *Edge //边列表，存储起点为Val的所有边
}

// Edge 边
type Edge struct {
	// 节点数组中的索引位置
	Idx    int
	Weight int
	Next   *Edge
}

// 求入度
func (g *Graph) GetPenetration(idx int) int {
	ret := 0
	for i := range g.Adj {
		if i != idx {
			next := g.Adj[i].FirstEdge
			for next != nil {
				if next.Idx == idx {
					ret++
				}
				next = next.Next
			}
		}
	}

	return ret
}

func (g *Graph) GetOut(idx int) int {
	ret := 0

	next := g.Adj[idx].FirstEdge
	for next != nil {
		ret++
		next = next.Next
	}

	return ret
}
