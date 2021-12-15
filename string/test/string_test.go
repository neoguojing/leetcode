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

func TestReverseVowels(t *testing.T) {

	s := "hello"
	ret := string.ReverseVowels(s)
	t.Log(ret)
}

func TestStrStr(t *testing.T) {

	ret := string.StrStr("hello", "ll")
	t.Log(ret)
}

func TestCountAndSay(t *testing.T) {

	ret := string.CountAndSay(5)
	t.Log(ret)
}

func TestNumDecodings(t *testing.T) {

	ret := string.NumDecodings("")
	t.Log(ret)
}
