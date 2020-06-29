package test

import (
	"leetcode/tree"
	"testing"
)

func TestNumTrees(t *testing.T) {
	ret := tree.NumTrees(3)
	t.Log(ret)
}

func TestGenerateBSTrees(t *testing.T) {
	ret := tree.GenerateBSTrees(3)
	for _, v := range ret {
		t.Log(v.ToString(v))
	}
}

func TestIsValidBST(t *testing.T) {
	root := tree.GeneBinaryTree([]int{2, 1, 3}, -1)
	ret := tree.IsValidBST(root)
	t.Log(ret)

	root = tree.GeneBinaryTree([]int{5, 1, 4, -1, -1, 3, 6}, -1)
	ret = tree.IsValidBST(root)
	t.Log(ret)
}
