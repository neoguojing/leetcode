package test

import (
	"leetcode/matrix"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	ret, count := matrix.SolveNQueens(4)
	t.Log(count)
	for _, row := range ret {
		for _, col := range row {
			t.Log(col)
		}
		t.Log("==============")
	}
}
