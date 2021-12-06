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

func TestCanReach(t *testing.T) {

	ret := array.CanReach("01101110", 2, 3)
	t.Log(ret)
	ret = array.CanReach("0000000000", 2, 5)
	t.Log(ret)

}
