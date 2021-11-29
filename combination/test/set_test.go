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

	in = [][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}
	ret = combination.Union(in)
	t.Log(ret)

	in = [][]int{{1, 4}, {0, 2}, {3, 5}}
	ret = combination.Union(in)
	t.Log(ret)
	t.Log(combination.HasUnion([]int{0, 4}, []int{3, 5}))

}
func TestInsert(t *testing.T) {
	in := [][]int{{1, 3}, {6, 9}}
	ret := combination.Insert(in, []int{2, 5})
	t.Log(ret)

	in = [][]int{{1, 5}}
	ret = combination.Insert(in, []int{6, 8})
	t.Log(ret)

	in = [][]int{{0, 5}, {9, 12}}
	ret = combination.Insert(in, []int{7, 16})
	t.Log(ret)
}

func TestRangeModule(t *testing.T) {
	rm := (combination.Constructor())
	(&rm).AddRange(10, 20)
	(&rm).RemoveRange(14, 16)
	t.Log((&rm).QueryRange(10, 14))
	t.Log((&rm).QueryRange(13, 15))
	t.Log((&rm).QueryRange(16, 17))
}

func TestCanPlaceFlowers(t *testing.T) {
	t.Log(combination.CanPlaceFlowers([]int{1, 0, 1, 0, 1, 0, 1}, 1))

}
