package test

import (
	"leetcode/sort"
	"testing"
)

func TestBucketSort(t *testing.T) {
	in := []int{6, 5, 3, 1, 8, 7, 2, 4}
	ret := sort.BucketSort(in, 5)
	t.Log(ret)
}
