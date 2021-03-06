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
	a := list.GeneListByArray([]int{1, 2, 3, 4, 5})

	head := list.RemoveNthFromEnd(a, 2)
	t.Log(head.ToString(head))
}

func TestReverseList(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 3, 4, 5})

	head, tail := list.ReverseList(a)
	t.Log(head.ToString(head))
	t.Log(tail)
}

func TestReverseKGroup(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 3, 4, 5})
	head := list.ReverseKGroup(a, 2)
	t.Log(head.ToString(head))
	a = list.GeneListByArray([]int{1, 2, 3, 4, 5})
	head = list.ReverseKGroup(a, 3)
	t.Log(head.ToString(head))
}

func TestRotateRight(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 3, 4, 5})
	head := list.RotateRight(a, 3)
	t.Log(head.ToString(head))

	a = list.GeneListByArray([]int{1, 2})
	head = list.RotateRight(a, 3)
	t.Log(head.ToString(head))

	a = list.GeneListByArray([]int{1, 2, 3})
	head = list.RotateRight(a, 3)
	t.Log(head.ToString(head))
}

func TestDeleteDuplicates(t *testing.T) {
	a := list.GeneListByArray([]int{1, 1, 2})
	head := list.DeleteDuplicates(a)
	t.Log(head.ToString(head))

	a = list.GeneListByArray([]int{1, 1, 2, 3, 3})
	head = list.DeleteDuplicates(a)
	t.Log(head.ToString(head))
}

func TestDeleteAllDuplicates(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 3, 3, 4, 4, 5})
	head := list.DeleteAllDuplicates(a)
	t.Log(head.ToString(head))

	a = list.GeneListByArray([]int{1, 1, 1, 2, 3})
	head = list.DeleteAllDuplicates(a)
	t.Log(head.ToString(head))
}

func TestPartition(t *testing.T) {
	a := list.GeneListByArray([]int{1, 4, 3, 2, 5, 2})
	head := list.Partition(a, 3)
	t.Log(head.ToString(head))
}

func TestReorderList(t *testing.T) {
	a := list.GeneListByArray([]int{1, 2, 3, 4, 5})
	head := list.ReorderList(a)
	t.Log(head.ToString(head))
}
