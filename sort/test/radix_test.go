package test

import (
	"leetcode/sort"
	"testing"
)

func TestRadixSort(t *testing.T) {
	in := []int{3, 6, 9, 1}
	in = sort.RadixSort(in)
	t.Log(in)

	in = []int{342, 58, 576, 356}
	in = sort.RadixSort(in)
	t.Log(in)
}
