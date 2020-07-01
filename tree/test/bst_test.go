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

func TestRecoverTree(t *testing.T) {
	root := tree.GeneBinaryTree([]int{1, 3, -1, -1, 2}, -1)
	t.Log(root.ToString(root))
	tree.RecoverTree(root)
	t.Log(root.ToString(root))

	root = tree.GeneBinaryTree([]int{3, 1, 4, -1, -1, 2}, -1)
	t.Log(root.ToString(root))
	tree.RecoverTree(root)
	t.Log(root.ToString(root))

}

func TestSortedArrayToBST(t *testing.T) {
	in := []int{-10, -3, 0, 5, 9}
	ret := tree.SortedArrayToBST(in)
	t.Log(ret.ToString(ret))
}
