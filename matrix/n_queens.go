package matrix

import (
	"leetcode/utils"
)

//SolveNQueens ...
// no 51
// no 52
// 摆皇后的位置，每行每列以及对角线只能出现 1 个皇后。输出所有的情况,和判定有几个解
// 回溯:递归深度是行，多路递归是列
func SolveNQueens(n int) ([][]string, int) {
	if n == 0 {
		return nil, 0
	}
	ret := make([][]string, 0)
	tmp := make([][]string, 0)
	count := solveNQueens(n, 0, tmp, &ret)

	return ret, count
}

func solveNQueens(n int, row int, tmp [][]string, ret *[][]string) int {
	if row == n {
		strRow := make([]string, 0)
		for _, r := range tmp {
			str := ""
			for _, c := range r {
				if c == "" {
					str += "."
				} else {
					str += c
				}
			}
			strRow = append(strRow, str)
		}
		*ret = append(*ret, strRow)
		return 1
	}

	count := 0
	rowTmp := make([]string, n)
	tmp = append(tmp, rowTmp)
	for i := 0; i < n; i++ {
		// 此处row非row+1
		if isColValid(row, i, tmp) {
			// 此处row非row+1
			if !isDiagonalValid(n, row, i, tmp) {
				continue
			} else {
				tmp[row][i] = "Q"
				// 注意row的选择
				count += solveNQueens(n, row+1, tmp, ret)
				//tmp = tmp[0 : len(tmp)-1]
				tmp[row][i] = ""
			}
		}
	}

	return count
}

// 判断列是否合法
func isColValid(row, col int, tmp [][]string) bool {
	for i := 0; i < row; i++ {
		if tmp[i][col] == "Q" {
			return false
		}
	}
	return true
}

// 判断对角线是否合法，列的差和行的差相等为对角线
func isDiagonalValid(n, row, col int, tmp [][]string) bool {

	for i := 0; i < row; i++ {
		for j := 0; j < n; j++ {
			if utils.Abs(i, row) == utils.Abs(j, col) {
				if tmp[i][j] == "Q" {
					return false
				}
			}
		}
	}
	return true
}
