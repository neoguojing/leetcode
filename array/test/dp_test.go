package test

import (
	"leetcode/array"
	"testing"
)

func TestPack01(t *testing.T) {
	t.Log(array.Pack01(8, []int{3, 5, 1, 2, 2}, []int{4, 5, 2, 1, 3}))
	t.Log(array.Pack01OP(8, []int{3, 5, 1, 2, 2}, []int{4, 5, 2, 1, 3}))
	t.Log(array.UnboundedPack(8, []int{1, 3, 4, 5}, []int{10, 40, 50, 70}))
}
