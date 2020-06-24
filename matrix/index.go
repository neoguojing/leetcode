package matrix

func initVisitedGraph(m, n int) [][]bool {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
		for j := range visited[i] {
			visited[i][j] = false
		}
	}

	return visited
}
