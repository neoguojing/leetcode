package string

//LongestPalindrome ...
//给定一个字符串，输出最长的回文子串。回文串指的是正的读和反的读是一样的字符串
func LongestPalindrome(in string) string {
	return longestPalindromeByDP1(in)
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

// todo
func longestPalindromeByDP1(in string) string {
	n := len(in)
	dp := make([]bool, n)
	for i := range dp {
		dp[i] = false
	}
	rs := []rune(in)
	maxLen := 0
	ret := ""
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			if in[i] == in[j] && (j-i < 3 || dp[j-1]) {
				dp[j] = true
			}

			if dp[j] && maxLen < j-i+1 {
				maxLen = j - i + 1
				ret = string(rs[i : j+1])
			}
		}
	}

	return ret
}
