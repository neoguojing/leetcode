package string

//LongestCommonPrefix ...
// no 14
// 最长公共前缀
// 递归法
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// 1.公共前缀和长度有关，只需要比较长度最短的字符串的个数
func LongestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	min := 1000000
	for i := range strs {
		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}

	if min == 0 {
		return ""
	}

	ret := ""
	for i := 0; i < min; i++ {
		flag := true
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] != strs[j][i] {
				flag = false
				break
			}
		}
		if flag {
			ret += string(strs[0][i])
		}
	}
	return ret
}
