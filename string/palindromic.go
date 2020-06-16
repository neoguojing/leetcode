package string

import()

//给定一个字符串，输出最长的回文子串。回文串指的是正的读和反的读是一样的字符串
func LongestPalindrome(in string) string {
	return longestPalindromeByDP(in)
}

//1.双重循环
//2.边循环边计算dp
/*
55
44 45  
33 34 35   
22 23 24 25  
11 12 13 14 15
*/
func longestPalindromeByDP(in string) string {
	n := len(in)
	dp := make([][]bool,n)
	for i,_:= range dp {
		dp[i] = make([]bool,0);
		for j:=0;j<n;j++ {
			dp[i] = append(dp[i],false)
		}
	}
	rs := []rune(in)
	maxLen := 0
	ret := ""
	for i:=n-1;i >= 0;i-- {
		for j:=i;j<n;j++{
			if in[i] == in[j] && i-j+1 <=2 {
				dp[i][j] = true
			} else if in[i] == in[j] && dp[i+1][j-1]{
				dp[i][j] = true
			}

			if dp[i][j]&& maxLen < i-j+1 {
				maxLen = i-j+1
				ret = string(rs[i:j+1])
			}
		}
	}

	return ret
}