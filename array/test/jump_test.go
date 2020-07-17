package test

import (
	"leetcode/array"
	"testing"
)

func TestCanJump(t *testing.T) {
	in := []int{2, 3, 1, 1, 4}
	ret := array.CanJump(in)
	t.Log(ret)

	in = []int{3, 2, 1, 0, 4}
	ret = array.CanJump(in)
	t.Log(ret)
}

func TestJump(t *testing.T) {
	in := []int{2, 3, 1, 1, 4}

	ret := array.Jump(in)
	t.Log(ret)

}
