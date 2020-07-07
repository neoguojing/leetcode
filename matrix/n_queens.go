package matrix

//SolveNQueens ...
// no 51
// 摆皇后的位置，每行每列以及对角线只能出现 1 个皇后。输出所有的情况
// 回溯法
func SolveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}

	return nil
}

func solveNQueens(n int, row int, col int, ret *[][]string) {
	if row == n {
		return
	}

	for i := 0; i < n; i++ {
		if i != col {
			solveNQueens(n, row+1, i, ret)
		}
	}
}

func isValid(n, row, col int) bool {
	return false
}
