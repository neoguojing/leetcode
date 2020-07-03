package test

import (
	"leetcode/combination"
	"testing"
)

func TestPremute(t *testing.T) {
	in := []int{1, 2, 3}
	ret := combination.Permute(in)
	t.Log(ret)
}

func TestPermuteUnique(t *testing.T) {
	in := []int{1, 1, 3}
	ret := combination.PermuteUnique(in)
	t.Log(ret)
}
