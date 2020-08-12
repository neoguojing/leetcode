package test

import (
	"leetcode/sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	in := []int{6, 5, 3, 1, 8, 7, 2, 4}
	sort.MergeSort(in, 0, len(in)-1)
	t.Log(in)
}

func TestQuickSort(t *testing.T) {
	in := []int{6, 5, 3, 1, 8, 7, 2, 4}
	sort.Quicksort(in, 0, len(in)-1)
	t.Log(in)
}

func TestBubbleSort(t *testing.T) {
	in := []int{6, 5, 3, 1, 8, 7, 2, 4}
	sort.BubbleSort(in)
	t.Log(in)
}

func TestHeapSort(t *testing.T) {
	in := []int{6, 5, 3, 1, 8, 7, 2, 4}
	sort.HeapSort(in)
	t.Log(in)
}
