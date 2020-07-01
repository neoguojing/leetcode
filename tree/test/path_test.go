package test

import (
	"leetcode/tree"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	in := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, -1, 1}
	root := tree.GeneBinaryTree(in, -1)
	ret := tree.HasPathSum(root, 22)
	t.Log(ret)
}

func TestPathSum(t *testing.T) {
	in := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1}
	root := tree.GeneBinaryTree(in, -1)
	ret := tree.PathSum(root, 22)
	t.Log(ret)
}
