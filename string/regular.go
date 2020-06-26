package string

/*IsMatch ...
1, If p[j] == s[i] :  dp[i][j] = dp[i-1][j-1];
2, If p[j] == '.' : dp[i][j] = dp[i-1][j-1];
3, If p[j] == '*':
			等于.的情况和尾字符相等的情况等价
               1   if p[j-1] != s[i] : dp[i][j] = dp[i][j-2]  //aab* aa  b*可以代表空字符串 ,则patten向前两个比较
               2   if p[j-1] == s[i] or p[j-1]== '.':
                              dp[i][j] = dp[i-1][j]    // ab* abbb   b*表示b出现大于1次的情况，则表示b*与text[i-1]
						   or dp[i][j] = dp[i][j-1]   // abb*   abb  b*表示出现1次，则从*前一位判断
                           or dp[i][j] = dp[i][j-2]   // abb*  ab    b*表示空，则patten 后退两位
*/
func IsMatch(text, patten string) bool {
	m := len(text)
	n := len(patten)
	dp := initDP2Ds(m+1, n+1)
	dp[0][0] = true
	//空字符串和出现*的字符比较，属于b*等于空的场景
	for i := 0; i < n; i++ {
		if patten[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text[i] == patten[j] || patten[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			}

			if patten[j] == '*' {
				if patten[j-1] != text[i] && patten[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j+1-2]

				} else if patten[j-1] == text[i] || patten[j-1] == '.' {
					dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j] || dp[i+1][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
