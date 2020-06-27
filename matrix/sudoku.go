package matrix

import "fmt"

//SudukuBoard ...
var SudukuBoard = [][]byte{
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
}

//InitSudukuBoard ...
func InitSudukuBoard() [][]byte {
	return nil
}

// IsValidSudoku ...
//no 36
/*
一个 9 * 9 的数独的棋盘。判断已经写入数字的棋盘是不是合法。需要满足下边三点，
每一行的数字不能重复
每一列的数字不能重复
9 个 3 * 3 的小棋盘中的数字也不能重复。
只能是 1 - 9 中的数字，不需要考虑数独最后能不能填满
*/
func IsValidSudoku(board [][]byte) bool {
	if board == nil {
		return false
	}
	rowMap := make(map[string]bool)

	for i := 0; i < len(board); i++ {
		for j := 0; i < len(board[0]); j++ {
			rowVal := fmt.Sprintf("row:%d:%d", i, board[i][j])
			if _, ok := rowMap[rowVal]; ok {
				return false
			} else {
				rowMap[rowVal] = true
			}
			colVal := fmt.Sprintf("col:%d:%d", j, board[i][j])
			if _, ok := rowMap[colVal]; ok {
				return false
			} else {
				rowMap[colVal] = true
			}

			cellVal := fmt.Sprintf("cell:%d:%d", whichCell(i, j), board[i][j])
			if _, ok := rowMap[cellVal]; ok {
				return false
			} else {
				rowMap[cellVal] = true
			}
		}
	}

	return true
}

func whichCell(row, col int) int {
	a := row / 3
	b := col / 3

	return b + a*3
}

//SolveSudoku ...
// no 37
// 给定一个数独棋盘，输出它的一个
// 回溯法：选项决定递归的广度(有多少个棵树),递归的深度,由子问题的个数界定,本题子问题的个数就是.的个数
//误区一: 首先棋盘填满才能判定是否合法?每一行是否合法,是由数字是否重复判定,与是否填满无关;
//优先判断要填入的数字合法,才会递归
//递归的深度是要填充位置的个数
// 可选择的值是1-9
//
func SolveSudoku(board [][]byte) [][]byte {
	solve(&board)
	return board
}

var input = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solve(board *[][]byte) bool {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (*board)[i][j] != '.' {
				continue
			}
			//可选项
			for k := range input {
				//先判断是否合法
				if isValidForPoint(i, j, *board, input[k]) {
					//填充
					(*board)[i][j] = input[k]
					//递归的填充下一个.成功则返回,否则回退
					if solve(board) {
						return true
					} else {
						(*board)[i][j] = '.'
					}
				}
			}
			return false
		}
	}

	return false
}

// 某个点的输入是否合法,行,列和cell,所谓合法是指该数字不重复,有没有.不重要
func isValidForPoint(x, y int, borad [][]byte, val byte) bool {
	//判断行是否合法
	for i := 0; i < 10; i++ {
		if borad[x][i] == val {
			return false
		}
	}

	for i := 0; i < 10; i++ {
		if borad[i][y] == val {
			return false
		}
	}

	left := y / 3 * 3
	right := left + 3
	top := x / 3 * 3
	bottom := top + 3

	for i := left; i < right; i++ {
		for j := top; j < bottom; j++ {
			if borad[i][j] == val {
				return false
			}
		}
	}
	return true
}
