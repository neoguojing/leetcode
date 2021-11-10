package array

import "fmt"

// no 518
// coins:找到合适的组合满足硬币需要
/*
amount = 5, coins = [1,2,5]
5=5

// 前半部分
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
1.1和2构成5的组合已知，
*/
func Change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1

	for i := 1; i <= len(coins); i++ {
		dp[i][0] = 1
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(coins)][amount]
}

// Pack01
// 01背包
func Pack01(tw int, w []int, v []int) int {
	dp := make([][]int, len(w)+1)
	for i := range dp {
		dp[i] = make([]int, tw+1)
		dp[i][0] = 0
	}
	useI := 0
	for i := 1; i <= len(w); i++ {
		for j := 1; j <= tw; j++ {
			if j-w[i-1] >= 0 {
				useI = dp[i-1][j-w[i-1]] + v[i-1]
			}
			if dp[i-1][j] > useI {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = useI
			}
		}
	}

	return dp[len(w)][tw]
}

// Pack01OP
// 01背包优化空间
// dp[i][j]的值只与dp[i-1][0,...,j-1]
// 使用滚动数组优化:逆向遍历，放置覆盖
func Pack01OP(tw int, w []int, v []int) int {
	dp := make([]int, tw+1)
	for i := 1; i <= len(w); i++ {
		for j := tw; j >= w[i-1]; j-- {
			if dp[j] > dp[j-w[i-1]]+v[i-1] {
				dp[j] = dp[j]
			} else {
				dp[j] = dp[j-w[i-1]] + v[i-1]
			}
			fmt.Println(dp)
		}
	}
	return dp[tw]
}

//UnboundedPack
//完全背包问题
// j必须正向遍历
func UnboundedPack(tw int, w []int, v []int) int {
	dp := make([]int, tw+1)
	for i := 1; i <= len(w); i++ {
		for j := w[i-1]; j <= tw; j++ {
			if dp[j] > dp[j-w[i-1]]+v[i-1] {
				dp[j] = dp[j]
			} else {
				dp[j] = dp[j-w[i-1]] + v[i-1]
			}
			fmt.Println(dp)
		}
	}
	return dp[tw]
}

// BoundPack
// 多重背包问题
// 空间优化j逆向遍历
func BoundPack(tw int, w []int, v []int, n []int) int {
	dp := make([]int, tw+1)

	for i := 1; i <= len(w); i++ {
		for j := tw; j >= w[i-1]; j-- {
			kn := j / w[i]
			if n[i] > j/w[i] {
				kn = n[i]
			}
			for k := 0; k <= kn; k++ {
				if dp[j] > dp[j-k*w[i-1]]+k*v[i-1] {
					dp[j] = dp[j]
				} else {
					dp[j] = dp[j-k*w[i-1]] + k*v[i-1]
				}
			}
		}
	}

	return dp[tw]
}
