package test

import (
	"leetcode/string"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	a := "aabcc"
	b := "dbbca"
	c := "aadbbcbcac"
	ret := string.IsInterleave(a, b, c)
	t.Log(ret)

	c = "aadbbbaccc"

	ret = string.IsInterleave(a, b, c)
	t.Log(ret)

}
