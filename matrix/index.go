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
			isFirstRowZero = true
			break
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			isFirstColZero = true
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

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
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

}

// Rotate 原地顺时针旋转矩阵90度
// no 48
// 转置加对称变换
func Rotate(matrix [][]int) {
	transpose(matrix)
	reflect(matrix)

}

func transpose(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[0]); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}

// 从左到右 （顺时针90度）
func reflect(matrix [][]int) {
	for j := 0; j < len(matrix[0])/2; j++ {
		for i := 0; i < len(matrix); i++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][len(matrix[0])-1-j]
			matrix[i][len(matrix[0])-1-j] = tmp
		}
	}
}

// FindRotation 确定矩阵是否可以通过旋转得到
// no 1886
// 坐标旋转
func FindRotation(mat [][]int, target [][]int) bool {
	var ret = [4]bool{true, true, true, true}
	n := len(mat)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != target[i][j] {
				ret[0] = false
			}

			if mat[i][j] != target[j][n-i-1] {
				ret[1] = false
			}

			if mat[i][j] != target[n-i-1][n-j-1] {
				ret[2] = false
			}

			if mat[i][j] != target[n-j-1][i] {
				ret[3] = false
			}

		}
	}

	return ret[0] || ret[1] || ret[2] || ret[3]
}

// no 54 矩阵的顺时针螺旋排序输出
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	ret := []int{}

	rEnd := len(matrix) - 1
	cEnd := len(matrix[0]) - 1
	rBegin, cBegin := 0, 0
	for rBegin <= rEnd && cBegin <= cEnd {
		//向右
		for j := cBegin; j <= cEnd; j++ {
			ret = append(ret, matrix[rBegin][j])
		}
		rBegin++
		//向下
		for i := rBegin; i <= rEnd; i++ {
			ret = append(ret, matrix[i][cEnd])
		}
		cEnd--
		//向左
		if rBegin <= rEnd {
			for j := cEnd; j >= cBegin; j-- {
				ret = append(ret, matrix[rEnd][j])
			}
			rEnd--
		}

		if cBegin <= cEnd {
			//向上
			for i := rEnd; i >= rBegin; i-- {
				ret = append(ret, matrix[i][cBegin])
			}
			cBegin++
		}
	}

	return ret
}
