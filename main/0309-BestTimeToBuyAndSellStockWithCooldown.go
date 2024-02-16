// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/16 17:23
// -----------------------------------------------
package main

// 动态规划

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([]int, 3)
	// 买入 卖出 冷冻期
	dp[0], dp[1], dp[2] = -prices[0], 0, 0
	for i := 1; i < n; i++ {
		dp[0], dp[1], dp[2] = max(dp[2]-prices[i], dp[0]), max(dp[0]+prices[i], dp[1]), max(dp[1], dp[2])
	}
	return max(dp[0], dp[1], dp[2])
}
