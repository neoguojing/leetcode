package string

import()

//「点.」代表任意字符，「星号*」 代表前一个字符重复 0 次或任意次
//dp[i][j] 代表以i开头到结尾的字符串和以j开头的字符串是否匹配
//. ab  a.    dp[i][j] = dp[i+1][j+1]
//* aa a*  ab a*b   dp[i][j] = dp[i+1][j+1]&&t[i-1]==t[i]||dp[i+1][j+1]&&t[i]==p[j+1]
func IsMatch(text,patten string) bool {
	m := len(text)
	n := len(patten)
	dp := initDP2Ds(m+1,n+1)
	dp[m][n] = true
	for i:=m-1;i>=0;i--{
		for j:=n-1;j>=0;j--{
			if text[i] == patten[j] && dp[i+1][j+1] {
				dp[i][j] = true
			} else if patten[j]==byte('.') && dp[i+1][j+1]{
				dp[i][j] = true
			} else{
				if (i > 0) {
					if patten[j]==byte('*') && dp[i+1][j+1] && (text[i]==text[i-1]||text[i]==patten[j+1]) {
						dp[i][j] = true
					}
				} 
			} 
		}
	}
	return dp[0][0]
}