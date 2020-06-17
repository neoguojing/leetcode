package string

/*IsMatch
1, If p[j] == s[i] :  dp[i][j] = dp[i-1][j-1];
2, If p[j] == '.' : dp[i][j] = dp[i-1][j-1];
3, If p[j] == '*':
               1   if p[j-1] != s[i] : dp[i][j] = dp[i][j-2]  //aab* aa  b*可以代表空字符串
               2   if p[j-1] == s[i] or p[j-1]== '.':
                              dp[i][j] = dp[i-1][j]    // ab* abbb
                           or dp[i][j] = dp[i][j-1]   // abb*   abb
                           or dp[i][j] = dp[i][j-2]   // abb*  ab
*/
func IsMatch(text, patten string) bool {
	m := len(text)
	n := len(patten)
	dp := initDP2Ds(m+1, n+1)
	dp[0][0] = true
	for i := 0; i < n; i++ {
		if patten[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

		}
	}
	return dp[m][n]
}
