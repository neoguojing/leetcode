package string

import (
	"leetcode/utils"
)

//LongestPalindrome ...
// no 5
//给定一个字符串，输出最长的回文子串。回文串指的是正的读和反的读是一样的字符串
func LongestPalindrome(in string) string {
	return longestPalindromeByDP(in)
}

//1.双重循环
//2.边循环边计算dp
/*
55
44 45
33 34 35                44
22 23 24 25             33  34
11 12 13 14 15          22  23  24
00 01 02 03 04 05       11  12  13 14
*/
func longestPalindromeByDP(in string) string {
	n := len(in)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, 0)
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], false)
		}
	}
	rs := []rune(in)
	maxLen := 0
	ret := ""
	//设计循环，i小于等于j必须的,且差值从0开始扩大
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if in[i] == in[j] && j-i+1 <= 2 {
				dp[i][j] = true
			} else if in[i] == in[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}

			if dp[i][j] && maxLen < j-i+1 {
				maxLen = j - i + 1
				ret = string(rs[i : j+1])
			}
		}
	}

	return ret
}

// LongestPalindromeSubseq ...
// no 516 求最长回文子序列 而不是子串
func LongestPalindromeSubseq(in string) int {
	if len(in) == 0 {
		return 0
	}
	n := len(in)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 0)
		for j := 0; j < n; j++ {
			if i == j {
				dp[i] = append(dp[i], 1)

			} else {
				dp[i] = append(dp[i], 0)
			}
		}
	}

	maxLen := 0
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if in[i] == in[j] {
				if j-i == 1 {
					dp[i][j] = 2
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = utils.Max(dp[i+1][j], dp[i][j-1])
			}

			if maxLen < dp[i][j] {
				maxLen = dp[i][j]
			}
		}
	}

	return maxLen
}

//CountSubstrings ...
// no 647
// 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
// dp[i][j] ==
func CountSubstrings(in string) int {

	if len(in) == 0 {
		return 0
	}

	if len(in) == 1 {
		return 1
	}

	n := len(in)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, 0)
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], false)
		}
	}
	count := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if in[i] == in[j] && i-j+1 <= 2 {
				count++
				dp[i][j] = true
			} else if in[i] == in[j] && dp[i+1][j-1] {
				dp[i][j] = true
				count++
			}
		}
	}

	return count
}
