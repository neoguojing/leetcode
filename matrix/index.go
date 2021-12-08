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

// SetZeroes
// no 73
// 给定一个矩阵，然后找到所有含有 0 的地方，把该位置所在行所在列的元素全部变成 0。
// 要求原地操作
func SetZeroes(matrix [][]int) {
	isFirstRowZero := false
	isFirstColZero := false
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			isFirstColZero = true
			break
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		if matrix[i][0] == 0 {
			isFirstRowZero = true
			break
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {

				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				if i != 0 {
					matrix[i][j] = 0
				}
			}
		}
	}

	if isFirstRowZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if isFirstColZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	return
}
