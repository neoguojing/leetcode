package graph

// Graph 邻接表
// 适合稀疏图
type Graph struct {
	E_num int
	V_num int
	Adj   []Vertex
}

// Vertex 节点
type Vertex struct {
	Val       int
	FirstEdge *Edge
}

// Edge 边
type Edge struct {
	// 节点数组中的索引位置
	Idx  int
	Next *Edge
}
