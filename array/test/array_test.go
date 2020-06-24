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
