package string

import (
	"sort"
)

//LengthOfLongestSubstring ...
//no 3 给定一个字符串，找到没有重复字符的最长子串，返回它的长度。
//滑动窗口
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	i, j := 0, 0
	hashMap := map[byte]int{}
	max := 0
	for i <= j && j < len(s) {
		hashMap[s[j]]++
		if v := hashMap[s[j]]; v > 1 {
			for i <= j {
				hashMap[s[i]]--
				if s[i] == s[j] {
					i++
					break
				}
				i++
			}
		}
		j++
		if j-i > max {
			max = j - i
		}
	}

	return max
}

// FindAnagrams 找到s中关于p的相同字母异序词 返回起始位置
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
	cnt := len(p)
	for j < len(s) {

		if set[s[j]-'a'] > 0 {
			cnt--
		}

		set[s[j]-'a']--
		j++

		if j-i == len(p) {
			if cnt == 0 {
				ret = append(ret, i)
			}

			if set[s[i]-'a'] >= 0 {
				cnt++
			}

			set[s[i]-'a']++
			i++
		}
	}
	return ret
}
