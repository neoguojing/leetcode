package array

import (
	"leetcode/utils"
	"math"
)

// MaxProfit1 ...
// no 121
//给一个数组，看作每天股票的价格，然后某一天买入，某一天卖出，最大收益可以是多少。可以不操作，收入就是 0。
//7,1,5,3,6,4
// 双指针
func MaxProfit1(prices []int) int {
	sell := 0
	buy := 0
	max := 0
	for ; sell < len(prices); sell++ {
		// 当价格变小是，保存为买入价
		if prices[sell] < prices[buy] {
			buy = sell
		} else {
			max = utils.Max(max, prices[sell]-prices[buy])
		}
	}
	return max
}

// MaxProfit2 ...
// no 122
// 可以不停的买入、卖出，但是只有卖出了才能继续买入,求最大收益
// 7,1,5,3,6,4
// 在跌之前卖出，之后买入，收益最大
func MaxProfit2(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

// MaxProfit3 ...
// no 123
// 给一个数组代表股票每天的价格。你最多可以买入卖出两次，但只有卖出了才可以再次买入，求出最大的收益是多少
// 状态机，每天的操作都有4种状态，买1，卖1，买2，卖2
func MaxProfit3(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	s1 := prices[0]
	s2, s3, s4 := math.MinInt32, math.MinInt32, math.MinInt32
	for i := 0; i < len(prices); i++ {
		s1 = utils.Max(s1, -prices[i])
		s2 = utils.Max(s2, s1+prices[i])
		s3 = utils.Max(s3, s2-prices[i])
		s4 = utils.Max(s4, s3+prices[i])
	}
	return utils.Max(0, s4)
}

// MaxProfit4 ...
// no 188
//给一个数组代表股票每天的价格。你最多可以买入卖出 K 次，但只有卖出了才可以再次买入，求出最大的收益是多少
// 动态规划：dp[i][k] 第i天最多交易k次的收益,k=0 dp[i][0] = 0,i = 0,dp[0][k] = 0  dp[i][k] = dp[i-1][k] 第i天什么都不做
// 要使第i天收益最大，则第i天选择卖出,则前面i-1天至少要有一天买入 dp[i][k] = max(dp[i-1][k],prices[i]-prices[j]+dp[j][k-1]),j =(0,i)
// 等价于：dp[i][k] = max(dp[i-1][k],prices[i]-（prices[j]-dp[j][k-1]）)
// 当k > len(prices)/2时，等价与任意买入，可参考122题
func MaxProfit4(K int, prices []int) int {
	if K == 0 || prices == nil || len(prices) == 0 {
		return 0
	}

	if K > len(prices)/2 {
		return MaxProfit2(prices)
	}

	dp := iniDP(len(prices), K+1, 0)

	for k := 1; k <= K; k++ {
		for i := 1; i < len(prices); i++ {
			min := math.MaxInt32
			for h := 0; h <= i; h++ {
				if prices[h]-dp[h][k-1] < min {
					min = prices[h] - dp[h][k-1]
				}
			}
			dp[i][k] = utils.Max(dp[i-1][k], prices[i]-min)
		}
	}

	return dp[len(prices)-1][K]
}
