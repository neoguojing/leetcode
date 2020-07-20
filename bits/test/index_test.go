package test

import (
	"leetcode/bits"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	t.Log(bits.HammingWeight(9))
}

func TestReverseBits(t *testing.T) {
	t.Logf("%b", 11)

	ret := bits.ReverseBits(11)
	t.Logf("%b", ret)
}

func TestAddBinary(t *testing.T) {

	ret := bits.AddBinary("1010", "1011")
	t.Log(ret)
}
