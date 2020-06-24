package string

func initDP2Ds(m, n int) [][]bool {
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, 0)
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], false)
		}
	}

	return dp
}

func initDP2DsForInt(m, n int) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, 0)
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], 0)
		}
	}

	return dp
}
