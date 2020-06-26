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
