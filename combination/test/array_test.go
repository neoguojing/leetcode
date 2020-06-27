package test

import (
	"leetcode/combination"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	ret := combination.NextPermutation([]int{1, 2, 3})
	t.Log(ret)

	ret = combination.NextPermutation([]int{3, 2, 1})
	t.Log(ret)

	ret = combination.NextPermutation([]int{1, 1, 6})
	t.Log(ret)
}

func TestCombinationSum(t *testing.T) {
	ret := combination.CombinationSum([]int{2, 3, 6, 7}, 7)
	t.Log(ret)
}
