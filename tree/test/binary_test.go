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
	ret := tree.GeneBinaryTree(nil, nil)
	t.Log(ret)
}
