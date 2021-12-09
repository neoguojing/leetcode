package string

import (
	"fmt"
	"sort"
	"strconv"
)

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

// StrStr 返回第一个出现的子串的起始位置
// no 28
func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(haystack) == 0 && len(needle) != 0 {
		fmt.Println("aaaaa")
		return -1
	}

	i, j := 0, 0
	for j < len(haystack) {
		if i >= len(needle) {
			break
		}

		if haystack[j] == needle[i] {
			i++
			j++
		} else {
			j = j - i + 1
			i = 0
		}
	}
	fmt.Println(i, j)
	if j >= len(haystack) && i != len(needle) {
		return -1
	}

	return j - len(needle)
}

func StrStrKMP(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(haystack) == 0 && len(needle) != 0 {
		fmt.Println("aaaaa")
		return -1
	}
	i, j := 0, 0
	next := fastNext(needle)
	for j < len(haystack) {
		if haystack[j] == needle[i] {
			j++
			i++
		} else if i != 0 {
			i = next[i-1]
		} else {
			j++
		}
		if i == len(needle) {
			break
		}
	}
	if j >= len(haystack) && i != len(needle) {
		return -1
	}
	return j - i
}

// getNext o(m*m)
func getNext(p string) []int {
	next := make([]int, len(p))
	next[0] = 0
	for i := 1; i < len(p); i++ {
		for k := 1; k <= i; k++ {
			if string(p[0:k]) == string(p[i-k+1:i+1]) {
				next[i] = k
			}
		}
	}

	return next
}

// fastNext
func fastNext(p string) []int {
	next := make([]int, len(p))
	next[0] = 0
	i := 1
	now := 0

	for i < len(p) {
		if p[now] == p[i] {
			next[i] = now + 1
			now++
			i++
		} else if now != 0 {
			now = next[now-1]
		} else {
			i++
		}
	}
	return next
}

// no 76
// 给两个字符串，S 和 T，在 S 中找出包含 T 中所有字母的最短字符串，不考虑顺序。
// todo

// no 91
// 每个数字对应一个字母，给一串数字，问有几种解码方式。例如 226 可以有三种，2|2|6，22|6，2|26。
// todo

// no 93
// todo

// no 38
func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	var ret string
	tmp := CountAndSay(n - 1)
	fmt.Println(tmp)
	count := 1
	i := 1
	for ; i < len(tmp); i++ {
		if tmp[i] == tmp[i-1] {
			count++
		} else {
			ret += strconv.Itoa(count) + string(tmp[i-1])
			count = 1
		}
	}

	if i == len(tmp) {
		ret += strconv.Itoa(count) + string(tmp[i-1])
	}

	return ret

}

// GroupAnagrams
// no 49 将由相同字母组成的单词分类
//  ["eat","tea","tan","ate","nat","bat"]
func GroupAnagrams(strs []string) [][]string {
	if strs == nil {
		return [][]string{}
	}
	set := map[string][]string{}
	ret := [][]string{}

	for _, str := range strs {
		letters := []string{}
		for _, letter := range str {
			letters = append(letters, string(letter))
		}
		sort.Strings(letters)
		sortStr := ""
		for _, str := range letters {
			sortStr += str
		}
		if _, ok := set[sortStr]; !ok {
			set[sortStr] = []string{}
		}
		set[sortStr] = append(set[sortStr], str)
	}

	for _, v := range set {
		ret = append(ret, v)
	}

	return ret
}

// IsAnagram
// no 242
func IsAnagram(s string, t string) bool {
	letters := []string{}
	for _, letter := range s {
		letters = append(letters, string(letter))
	}
	sort.Strings(letters)
	s1 := ""
	for _, str := range letters {
		s1 += str
	}
	letters1 := []string{}
	for _, letter := range t {
		letters1 = append(letters1, string(letter))
	}
	sort.Strings(letters1)
	t1 := ""
	for _, str := range letters1 {
		t1 += str
	}

	return s1 == t1
}
