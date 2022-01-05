package string

// WordBreak
// 139 s是否能够完全分裂为wordDict中的单词；s不能有剩余的非dict单词
// 递归超时
func WordBreak(s string, wordDict []string) bool {
	set := map[string]bool{}
	for _, w := range wordDict {
		set[w] = true
	}
	return wordBreakHelper(s, set)
}

func wordBreakHelper(s string, set map[string]bool) bool {
	if len(s) == 0 {
		return true
	}

	for i := 1; i < len(s); i++ {
		if _, ok := set[s[0:i]]; ok && wordBreakHelper(s[i:], set) {
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
	dp[0] = true
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

// WordBreakII 将s分解为带空格的句子，句子中每个单词都是wordDict中的单词
// 140
func WordBreakII(s string, wordDict []string) []string {
	set := map[string]bool{}
	for _, w := range wordDict {
		set[w] = true
	}

	dp := make([][]string, len(s)+1)
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := set[s[j:i]]; ok && (j == 0 || dp[j] != nil) {
				if dp[j] != nil {
					for k := range dp[j] {
						if dp[i] == nil {
							dp[i] = make([]string, 1)
							dp[i][0] = dp[j][k] + " " + s[j:i]
						} else {
							dp[i] = append(dp[i], dp[j][k]+" "+s[j:i])
						}

					}
				}

				if j == 0 {
					if dp[i] == nil {
						dp[i] = []string{}
					}
					dp[i] = append(dp[i], s[j:i])
				}
			}
		}
	}

	return dp[len(s)]
}
