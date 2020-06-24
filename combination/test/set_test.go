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

	in = in[0:0]
	t.Log(in)
}
