package test

import (
	strings "leetcode/string"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	in := "babad"

	t.Log(strings.LongestPalindrome(in))
	in = "cbbd"
	t.Log(strings.LongestPalindrome(in))
}

func TestCountSubstrings(t *testing.T) {
	in := "abc"

	t.Log(strings.CountSubstrings(in))
	in = "aaa"
	t.Log(strings.CountSubstrings(in))

}

func TestLongestPalindromeSubseq(t *testing.T) {
	in := "bbbab"

	t.Log(strings.LongestPalindromeSubseq(in))
	in = "cbbd"
	t.Log(strings.LongestPalindromeSubseq(in))
	in = "abcabcabcabc"
	t.Log(strings.LongestPalindromeSubseq(in))
}
