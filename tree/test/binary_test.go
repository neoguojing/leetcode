package test

import (
	"leetcode/tree"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	a := tree.GeneBinaryTree([]int{1, 2}, 0)
	b := tree.GeneBinaryTree([]int{1, 0, 2}, 0)
	a.Print(a)
	b.Print(b)
	ret := tree.IsSameTree(a, b)
	t.Log(ret)

	a = tree.GeneBinaryTree([]int{1, 2, 1}, 0)
	b = tree.GeneBinaryTree([]int{1, 1, 2}, 0)
	a.Print(a)
	b.Print(b)
	ret = tree.IsSameTree(a, b)
	t.Log(ret)
}

func TestIsSymmetric(t *testing.T) {
	a := tree.GeneBinaryTree([]int{1, 2, 2, 3, 4, 4, 3}, 0)
	a.Print(a)
	ret := tree.IsSymmetric(a)
	t.Log(ret)

	a = tree.GeneBinaryTree([]int{1, 2, 2, 0, 3, 0, 3}, 0)
	a.Print(a)
	ret = tree.IsSymmetric(a)
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
		ret = append(ret, node.Val.(int))
	}
	tree.InOrder(root, op)
	t.Log(ret)
}

func TestBFS(t *testing.T) {
	in := []int{1, -1, 2, -1, -1, 3}
	root := tree.GeneBinaryTree(in, -1)
	ret := tree.BFS(root, nil)
	t.Log(ret)
}
func TestBFSByRow(t *testing.T) {
	in := []int{1, -1, 2, -1, -1, 3}
	root := tree.GeneBinaryTree(in, -1)
	t.Log(root.ToString(root))
	dep := tree.MaxDepth(root)
	t.Log(dep)
	ret := tree.BFSByRow(root)
	t.Log(ret)
	t.Log(root.ToString(root))

	in = []int{1, 2, 3, 4, 5, 6, 7, 1, 1, 1, 1, 1, 1, 1, 1}
	root = tree.GeneBinaryTree(in, -1)
	t.Log(root.ToString(root))
}

func TestBuildTree(t *testing.T) {
	preorder := []int{1, 2}
	inorder := []int{1, 2}
	ret := tree.BuildTree(preorder, inorder)
	t.Log(ret.ToString(ret))
}

func TestBuildTreeByPost(t *testing.T) {
	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	ret := tree.BuildTreeByPost(postorder, inorder)
	t.Log(ret.ToString(ret))
}

func TestMinDepth(t *testing.T) {
	in := []int{3, 9, 20, -1, -1, 15, 7}
	root := tree.GeneBinaryTree(in, -1)
	ret := tree.MinDepth(root)
	t.Log(ret)
}

func TestFlatten(t *testing.T) {
	in := []int{1, 2, 5, 3, 4, -1, 6}
	root := tree.GeneBinaryTree(in, -1)
	tree.Flatten(root)
	t.Log(root.ToString(root))
}
