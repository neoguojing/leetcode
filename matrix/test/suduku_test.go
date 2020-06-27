package test

import (
	"leetcode/matrix"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {

	ret := matrix.IsValidSudoku(matrix.SudukuBoard)
	if ret == false {
		t.Error(ret)
	}

	t.Log(8 / 3 * 3)
}
