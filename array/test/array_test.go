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

func TestSearchInsert(t *testing.T) {
	in := []int{1, 3, 5, 6}
	ret := array.SearchInsert(in, 5)
	t.Log(ret)

	ret = array.SearchInsert(in, 2)
	t.Log(ret)

	ret = array.SearchInsert(in, 7)
	t.Log(ret)

	ret = array.SearchInsert(in, 0)
	t.Log(ret)

}

func TestRemoveElement(t *testing.T) {
	in := []int{3, 2, 2, 3}
	ret := array.RemoveElement(in, 3)
	t.Log(ret)

	in = []int{0, 1, 2, 2, 3, 0, 4, 2}
	ret = array.RemoveElement(in, 2)
	t.Log(ret)
}

func TestSearchInRotateArray(t *testing.T) {

	ret := array.SearchInRotateArray([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8)
	t.Log(ret)
	ret = array.SearchInRotateArray([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	t.Log(ret)
	ret = array.SearchInRotateArray([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	t.Log(ret)
	ret = array.SearchInRotateArray([]int{1}, 0)
	t.Log(ret)

	ret1 := array.SearchInRotateArray2([]int{1, 0, 1, 1, 1}, 0)
	t.Log(ret1)
	ret1 = array.SearchInRotateArray2([]int{1, 2, 1, 1, 1}, 2)
	t.Log(ret1)

}
