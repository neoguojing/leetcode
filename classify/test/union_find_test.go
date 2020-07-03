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
