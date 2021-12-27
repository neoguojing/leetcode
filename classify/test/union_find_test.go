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

func TestLongestConsecutive(t *testing.T) {
	grid := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	ret := classify.LongestConsecutive(grid)
	t.Log(ret)
	grid = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	ret = classify.LongestConsecutive(grid)
	t.Log(ret)
}
