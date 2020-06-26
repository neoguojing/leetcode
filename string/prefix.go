package string

//LongestCommonPrefix ...
// no 14
// 最长公共前缀
// 递归法
func LongestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 || len(strs) == 1 {
		return ""
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
			}
		}
		if flag {
			ret += string(strs[0][i])
		}
	}
	return ret
}
