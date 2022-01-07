package test

import (
	"leetcode/tree"
	"testing"
)

func TestBiTree(t *testing.T) {
	bit := tree.NewBiTree([]int{1, 2, 3, 4, 5})
	t.Log(bit.GetSum(0))
	t.Log(bit)
}
