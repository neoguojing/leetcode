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

func TestTrap(t *testing.T) {
	in := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := array.Trap(in)
	t.Log(ret)
}
