package test

import (
	"leetcode/matrix"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	ret := matrix.UniquePaths(3, 2)
	t.Log(ret)
	ret = matrix.UniquePaths(7, 3)
	t.Log(ret)
}

func TestUniquePathsWithObstacles(t *testing.T) {
	in := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	ret := matrix.UniquePathsWithObstacles(in)
	t.Log(ret)
}

func TestMinPathSum(t *testing.T) {
	in := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	ret := matrix.MinPathSum(in)
	t.Log(ret)
}
