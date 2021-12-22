package string

import "fmt"

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
func longestPalindromeByDP(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	max := 0
	var ret string
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if j == i {
				dp[i][j] = true
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}

			if dp[i][j] && max < j-i+1 {
				max = j - i + 1
				ret = s[i : j+1]
			}
		}
	}
	return ret
}

// LongestPalindromeSubseq ...
// no 516 求最长回文子序列 而不是子串
func LongestPalindromeSubseq(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j {
				dp[i][j] = 1
			} else if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				if dp[i+1][j] > dp[i][j-1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}

		}

	}
	fmt.Println(dp)
	return dp[0][len(s)-1]
}

//CountSubstrings ...
// no 647
// 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
// dp[i][j] ==
func CountSubstrings(s string) int {

	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if j == i {
				dp[i][j] = true
				count++
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
				count++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				count++
			}
		}
	}
	return count
}

// ShortestPalindrome 在前面添加字符，使新的字符串为回文；且字符串最短
// no 214
// aacecaaa
func ShortestPalindrome(s string) string {
	rev := ""
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}

	for i := 0; i < len(s); i++ {
		if string(s[0:len(s)-i]) == rev[i:] {
			return string(rev[0:i]) + s
		}
	}
	return ""
}

// IsPalindrome
// 125
func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] < '0' || (s[i] < 'A' && s[i] > '9') || (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z' {
			i++
			continue
		}

		if s[j] < '0' || (s[j] < 'A' && s[j] > '9') || (s[j] > 'Z' && s[j] < 'a') || s[j] > 'z' {
			j--
			continue
		}

		a := s[i]
		if a < 'a' {
			a = s[i] + ('a' - 'A')
		}
		b := s[j]
		if b < 'a' {
			b = s[j] + ('a' - 'A')
		}

		if a != b {
			return false
		}
		i++
		j--
	}

	return true
}

// ValidPalindrome 去掉最多一个字符能否构成回文
// 680
// abca
// aguokepatgbnvfqmg m l cupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupucu l m gmqfvnbgtapekouga
func ValidPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return isPalindrome(s, i+1, j) || isPalindrome(s, i, j-1)
		}
	}

	return true
}

func isPalindrome(s string, l, r int) bool {

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

//Partition 将字符串分组，使得所有的分组都是回文，返回所有可能
// 131
// 回溯 + dp
func Partition(s string) [][]string {

	ret := [][]string{}
	backward(0, s, []string{}, &ret)
	return ret
}

func backward(idx int, s string, one []string, ret *[][]string) {
	if idx >= len(s) {
		*ret = append(*ret, one)
	}

	for i := range s[idx:] {
		if isPalindrome(s, 0, i) {
			one = append(one, s[0:i])
			backward(i+1, s, one, ret)
			one = one[:len(one)-1]
		}
	}
}
