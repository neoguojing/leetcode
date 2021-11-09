package string

import "fmt"

// IsInterleave ...
// no 97
// 在两个字符串 s1 和 s2 中依次取字母，问是否可以组成 S3。什么意思呢？比如 s1 = abc , s2 = de，s3 = abdce。
// len(s1) + len(s2) == len(s3)
//回溯
func IsInterleave(s1, s2, s3 string) bool {
	return isInterleave(s1, s2, s3, 0, 0, 0)
}

func isInterleave(s1, s2, s3 string, i, j, k int) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if i == len(s1) && j == len(s2) && k == len(s3) {
		return true
	}

	if i == len(s1) {
		for j < len(s2) {
			if s2[j] != s3[k] {
				return false
			}
			j++
			k++
		}
		return true
	}

	if j == len(s2) {
		for i < len(s1) {
			if s1[i] != s3[k] {
				return false
			}
			i++
			k++
		}
		return true
	}

	if s1[i] == s3[k] {
		if isInterleave(s1, s2, s3, i+1, j, k+1) {
			return true
		}
	}

	if s2[j] == s3[k] {
		if isInterleave(s1, s2, s3, i, j+1, k+1) {
			return true
		}
	}

	return false
}

// ReverseStr ...
// no 541
// Input: s = "abcdefg", k = 2
// Output: "bacdfeg"
// 开头k个字符反转，同时每隔2k个字符反转依次，每次反转k个字符
// 等价于：每2k字符，反转前k个
func ReverseStr(s string, k int) string {
	if k > len(s) {
		return ReverseString([]byte(s))
	}

	if len(s) >= k && len(s) < 2*k {
		return ReverseString([]byte(s[:k])) + s[k:]
	}
	var ret string
	ret = ReverseString([]byte(s[:k])) + s[k:2*k]
	ret += ReverseStr(s[2*k:], k)

	return ret
}

// ReverseString ...
// no 344
// 原地反转string
func ReverseString(s []byte) string {
	if s == nil || len(s) == 1 || len(s) == 0 {
		return string(s)
	}

	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = tmp
	}
	return string(s)
}

// ReverseVowels
// no 345
// 交换元音字符
func ReverseVowels(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	vowels := func(in byte) bool {
		switch in {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}
	ret := []byte(s)
	i, j := 0, len(s)-1
	fmt.Println(i, j)
	for i < j {

		for i < j && !vowels(s[i]) {
			i++
		}
		for i < j && !vowels(s[j]) {
			j--
		}

		if vowels(s[i]) {
			if vowels(s[j]) {
				tmp := ret[i]
				ret[i] = ret[j]
				ret[j] = tmp
			}
		}
		i++
		j--

	}
	return string(ret)
}

// no 76
// 给两个字符串，S 和 T，在 S 中找出包含 T 中所有字母的最短字符串，不考虑顺序。
// todo

// no 91
// 每个数字对应一个字母，给一串数字，问有几种解码方式。例如 226 可以有三种，2|2|6，22|6，2|26。
// todo

// no 93
// todo
