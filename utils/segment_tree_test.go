package utils

import "testing"

func TestSegTree(t *testing.T) {
	arr := []int{18, 17, 13, 19, 15, 11, 20, 12, 33, 25}

	merge := func(i, j int) int {
		return i + j
	}
	seg := NewSegmentTree(arr, merge)

	t.Log(seg)

	q := seg.Query(2, 8, merge)
	if q != 123 {
		t.Error(q)
	}

	seg.Update(1, 3, merge)
	t.Log(seg)
}
