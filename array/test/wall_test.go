package test

import (
	"leetcode/array"
	"testing"
)

func TestMaxArea(t *testing.T) {
	in := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret := array.MaxArea(in)
	t.Log(ret)
}
