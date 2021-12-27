package string

// WordBreak
// 139 s是否能够完全分裂为wordDict中的单词；s不能有剩余的非dict单词
// 递归超时
func WordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for _, word := range wordDict {
		if len(word) <= len(s) && s[:len(word)] == word {
			ret := WordBreak(s[len(word):], wordDict)
			if ret {
				return true
			}
		}
	}
	return false
}
