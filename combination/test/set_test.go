package test

import (
	"leetcode/combination"
	"testing"
)

func TestSubsets(t *testing.T) {
	in := []int{1, 2, 3}
	ret := combination.Subsets(in)
	t.Log(ret)

}
