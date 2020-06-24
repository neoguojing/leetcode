package test

import (
	"leetcode/matrix"
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]rune{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	word := "ABCCED"
	ret := matrix.Exist(board, word)
	if ret == false {
		t.Error(ret)
	}

	word = "SEE"
	ret = matrix.Exist(board, word)
	if ret == false {
		t.Error(ret)
	}

	word = "ABCB"
	ret = matrix.Exist(board, word)
	if ret == true {
		t.Error(ret)
	}
}
