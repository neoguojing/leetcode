package test

import (
	"leetcode/integer"
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	t.Log(integer.NthUglyNumber(10))

	t.Log(integer.NthUglyNumber1(5, 2, 11, 13))

	t.Log(integer.NthSuperUglyNumber(12, []int{2, 7, 13, 19}))
}
