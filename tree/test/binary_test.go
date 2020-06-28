package test

import (
	"leetcode/tree"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	ret := tree.IsSameTree(nil, nil)
	t.Log(ret)
}

func TestGeneBinaryTree(t *testing.T) {
	in := []int{1, -1, 2, -1, -1, 3}
	ret := tree.GeneBinaryTree(in, -1)
	t.Log(ret.ToString(ret))
}
