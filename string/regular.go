package string

import()

//「点.」代表任意字符，「星号*」 代表前一个字符重复 0 次或任意次
//. ab  a.    dp[i][j] = dp[i-1][j-1]
//* aa a*  ab a*b   dp[i][j] = dp[i-1][j-1]&&t[i-1]==t[i]||dp[i-1][j-1]&&t[i]==p[j+1]
func isMatch(text,patten string) bool {
	for i:=0;i<len(text);i++{
		for j:=0;j<len(patten);j++{
			
		}
	}
	return false
}