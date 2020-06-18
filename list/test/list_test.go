package test

import (
	"leetcode/list"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 3, 4})

	b := list.SwapPairs(a)
	t.Log(b.ToString(b))
}
