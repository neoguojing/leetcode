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
