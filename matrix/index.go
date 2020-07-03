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

func iniDP(m, n int, tag int) [][]int {
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
		for j := range visited[i] {
			visited[i][j] = tag
		}
	}

	return visited
}
