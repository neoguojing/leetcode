package test

import (
	"leetcode/classify"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]int{
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{0, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 1},
	}
	ret := classify.NumIslands(grid)
	t.Log(ret)
}

func TestSolve(t *testing.T) {
	grid := [][]string{
		[]string{"X", "X", "X", "X"},
		[]string{"X", "O", "O", "X"},
		[]string{"X", "X", "O", "X"},
		[]string{"X", "O", "X", "X"},
	}
	ret := classify.Solve(grid)
	t.Log(ret)
}
