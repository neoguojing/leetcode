package test

import (
	"leetcode/list"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 3, 4})

	b := list.SwapPairs(a)
	t.Log(b.ToString(b))
}

func TestRemoveNthFromEnd(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 3, 4,5})

	head := list.RemoveNthFromEnd(a,2)
	t.Log(head.ToString(head))
}

func TestRReverseList(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 3, 4,5})

	head := list.ReverseList(a)
	t.Log(head.ToString(head))
}