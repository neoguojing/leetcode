package test

import (
	"leetcode/array"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	in := []int{1, 1, 2, 2, 3, 3}
	ret := array.RemoveDuplicates(in)
	t.Log(ret)
	t.Log(in[0:ret])
}

func TestSearchRange(t *testing.T) {
	in := []int{5, 7, 7, 8, 8, 10}
	ret := array.SearchRange(in, 8)
	t.Log(ret)

	ret = array.SearchRange(in, 6)
	t.Log(ret)

}
