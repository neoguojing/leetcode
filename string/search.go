package string

import (
	"sort"
)

//LengthOfLongestSubstring ...
//no 3给定一个字符串，找到没有重复字符的最长子串，返回它的长度。
//滑动窗口
func LengthOfLongestSubstring(in string) int {
	max := 0
	set := make(map[byte]int)
	for start, end := 0, 0; end < len(in); end++ {
		if _, ok := set[in[end]]; ok {
			if max < (end - start) {
				max = end - start
			}
			delete(set, in[start])
			start++
		}
		set[in[end]] = 1
	}

	return max
}

// FindAnagrams 找到s中关于p的相同字母异序词
// 438
func FindAnagrams(s string, p string) []int {
	bytes := []byte(p)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	p = string(bytes)
	ret := []int{}
	i, j := 0, len(p)-1
	for ; j < len(s); i, j = i+1, j+1 {
		tmp := []byte(s[i : j+1])
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		if string(tmp) == p {
			ret = append(ret, i)
		}
	}
	return ret
}

func FindAnagramsWithSlideWindow(s string, p string) []int {
	set := [26]int{}
	for i := range p {
		set[p[i]-'a']++
	}

	ret := []int{}
	i, j := 0, 0
	for i <= j && j < len(s) {
		if set[s[i]-'a'] == 0 {
			i++
			j++
			continue
		}

		set[s[i]-'a']--

	}
	return ret
}
