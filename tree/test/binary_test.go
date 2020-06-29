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

func TestInOrder(t *testing.T) {
	in := []int{1, -1, 2, -1, -1, 3}
	root := tree.GeneBinaryTree(in, -1)

	ret := make([]int, 0)
	op := func(node *tree.TreeNode) {
		ret = append(ret, node.Val)
	}
	tree.InOrder(root, op)
	t.Log(ret)
}

func TestBFSByRow(t *testing.T) {
	in := []int{1, -1, 2, -1, -1, 3}
	root := tree.GeneBinaryTree(in, -1)
	t.Log(root.ToString(root))
	dep := tree.Depth(root)
	t.Log(dep)
	ret := tree.BFSByRow(root)
	t.Log(ret)
	t.Log(root.ToString(root))
}
