package matrix

import "leetcode/utils"

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
			matrix[i][j], matrix[i][len(matrix[0])-1-j] = matrix[i][len(matrix[0])-1-j], matrix[i][j]
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
	dir := 0
	for rBegin <= rEnd && cBegin <= cEnd {
		//向右
		if dir == 0 {

			for j := cBegin; j <= cEnd; j++ {
				ret = append(ret, matrix[rBegin][j])
			}
			rBegin++
			dir = 1
		} else if dir == 1 {
			for i := rBegin; i <= rEnd; i++ {
				ret = append(ret, matrix[i][cEnd])
			}
			cEnd--
			dir = 2
		} else if dir == 2 {
			for j := cEnd; j >= cBegin; j-- {
				ret = append(ret, matrix[rEnd][j])
			}
			rEnd--
			dir = 3
		} else if dir == 3 {
			//向上
			for i := rEnd; i >= rBegin; i-- {
				ret = append(ret, matrix[i][cBegin])
			}
			cBegin++
			dir = 0
		}
	}

	return ret
}

// GenerateMatrix 1-n平方的数填到n*n的矩阵里
// no 59
func GenerateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{{1}}
	}
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	rs, re, cs, ce := 0, n-1, 0, n-1
	dir := 0
	i := 0
	for rs <= re && cs <= ce {
		switch dir {
		case 0:
			for c := cs; c <= ce; c++ {
				matrix[rs][c] = i
				i++
			}
			rs++
			dir = 1
		case 1:
			for r := rs; r <= re; r++ {
				matrix[r][ce] = i
				i++
			}
			ce--
			dir = 2
		case 2:
			for c := ce; c >= cs; c-- {
				matrix[re][c] = i
				i++
			}
			re--
			dir = 3
		case 3:
			for r := re; r >= rs; r-- {
				matrix[r][cs] = i
				i++
			}
			cs++
			dir = 0
		}
	}

	return matrix
}

// SpiralMatrixIII 从起点顺时针螺旋遍历矩阵；步长按圈，每次半径加1
// no 885
func SpiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	ret := make([][]int, rows*cols)
	ret[0] = []int{rStart, cStart}

	dir := 0
	rb, re, cb, ce := rStart-1, rStart+1, cStart-1, cStart+1
	i := 1
	for rb >= 0 || re < rows || cb >= 0 || ce < cols {
		switch dir {
		case 0: //向下
			for r := rStart; r <= re; r++ {
				if r >= 0 && r < rows && ce >= 0 && ce < cols {
					ret[i] = []int{r, ce}
					i++
				}
			}
			dir = 1
		case 1: //向左
			for c := ce - 1; c >= cb; c-- {
				if c >= 0 && c < cols && re >= 0 && re < rows {
					ret[i] = []int{re, c}
					i++
				}
			}
			dir = 2
		case 2: //向上
			for r := re - 1; r >= rb; r-- {
				if r >= 0 && r < rows && cb >= 0 && cb < cols {
					ret[i] = []int{r, cb}
					i++
				}
			}
			dir = 3
		case 3: //向右
			for c := cb + 1; c <= ce; c++ {
				if c >= 0 && c < cols && rb >= 0 && rb < rows {
					ret[i] = []int{rb, c}
					i++
				}
			}

			dir = 0
			rStart = rb
			rb--
			re++
			cb--
			ce++
		}
	}

	return ret
}

// MaximalSquare 由1组成的最大正方形面积
// 221
// dp[i][j] 表示以(i,j)为右下角的正方形的的边长
// dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
func MaximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	maxSide := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = utils.Min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}
