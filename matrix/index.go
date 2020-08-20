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

// 73
// 给定一个矩阵，然后找到所有含有 0 的地方，把该位置所在行所在列的元素全部变成 0。
// 要求原地place
