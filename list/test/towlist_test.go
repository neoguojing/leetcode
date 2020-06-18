package test

import (
	"leetcode/list"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 4})
	b := list.GeneListByArray([]int{1, 3, 4})
	c := list.MergeTwoLists(a, b)

	t.Log(list.ToString(c))
}
