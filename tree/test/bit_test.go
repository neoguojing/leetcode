package test

import (
	"leetcode/tree"
	"testing"
)

func TestBiTree(t *testing.T) {
	bit := tree.NewBiTree([]int{1, 2, 3, 4, 5})

	t.Log(bit.GetSum(3))
	t.Log(bit.RangeSum(2, 4))
	t.Log(bit.Get(4))
	t.Log(bit.GetArray())
	bit.Set(4, 6)
	t.Log(bit.RangeSum(2, 4))
	t.Log(bit.Get(4))
	t.Log(bit.GetArray())
}
