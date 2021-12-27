package string

// WordBreak
// 139 s是否能够完全分裂为wordDict中的单词；s不能有剩余的非dict单词
// 递归超时
func WordBreak(s string, wordDict []string) bool {
	set := map[string]bool{}
	for _, w := range wordDict {
		set[w] = true
	}
	mem := map[string]bool{}
	return wordBreakHelper(s, set, mem)
}

func wordBreakHelper(s string, set, mem map[string]bool) bool {
	if len(s) == 0 {
		return true
	}

	for i := 1; i < len(s); i++ {
		if _, ok := set[s[0:i]]; ok && wordBreakHelper(s[i:], set, mem) {
			return true
		}
	}
	return false
}

// dp[i] = true
// dp[i] = dp[i-j]  && dp[]
func WordBreakDP(s string, wordDict []string) bool {
	set := map[string]bool{}
	for _, w := range wordDict {
		set[w] = true
	}
	dp := make([]bool, len(s)+1)
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := set[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}

	}
	return dp[len(s)]
}
