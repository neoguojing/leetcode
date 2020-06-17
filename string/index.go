package string

import()

func initDP2Ds(m,n int) [][]bool {
	dp := make([][]bool, m)
	for i, _ := range dp {
		dp[i] = make([]bool, 0)
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], false)
		}
	}

	return dp
}