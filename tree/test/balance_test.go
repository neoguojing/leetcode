package test

import (
	"leetcode/tree"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	in := []int{3, 9, 20, -1, -1, 15, 7}
	root := tree.GeneBinaryTree(in, -1)
	ret := tree.IsBalanced(root)
	t.Log(ret)

	in = []int{1, 2, 2, 3, 3, -1, -1, 4, 4}
	root = tree.GeneBinaryTree(in, -1)
	ret = tree.IsBalanced(root)
	t.Log(ret)

}
