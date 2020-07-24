package test

import (
	"leetcode/array"
	"testing"
)

func TestMaxProfit1(t *testing.T) {
	in := []int{7, 1, 5, 3, 6, 4}
	ret := array.MaxProfit1(in)
	t.Log(ret)

	in = []int{7, 6, 4, 3, 1}
	ret = array.MaxProfit1(in)
	t.Log(ret)
}

func TestMaxProfit2(t *testing.T) {
	in := []int{7, 1, 5, 3, 6, 4}
	ret := array.MaxProfit2(in)
	t.Log(ret)
}

func TestMaxProfit3(t *testing.T) {
	in := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := array.Trap(in)
	t.Log(ret)
}

func TestMaxProfit4(t *testing.T) {
	in := []int{3, 3, 5, 0, 0, 3, 1, 4}
	ret := array.MaxProfit4(2, in)
	t.Log(ret)

	in = []int{3, 2, 6, 5, 0, 3}
	ret = array.MaxProfit4(2, in)
	t.Log(ret)
}
