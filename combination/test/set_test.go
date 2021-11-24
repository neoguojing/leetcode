package test

import (
	"leetcode/combination"
	"testing"
)

func TestSubsets(t *testing.T) {
	in := []int{1, 2, 3}
	ret := combination.Subsets(in)
	t.Log(ret)

}
func TestSubsetsWithBit(t *testing.T) {
	in := []int{1, 2, 3}
	ret := combination.SubsetsWithBit(in)
	t.Log(ret)

}

func TestSubsetsWithDup(t *testing.T) {
	in := []int{1, 2, 2}
	ret := combination.SubsetsWithDup(in)
	t.Log(ret)
}

func TestUnion(t *testing.T) {
	in := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	ret := combination.Union(in)
	t.Log(ret)
}
