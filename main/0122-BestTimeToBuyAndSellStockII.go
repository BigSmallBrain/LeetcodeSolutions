// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/14 18:16
// -----------------------------------------------
package main

// 动态规划

func maxProfit0122(prices []int) int {
	n := len(prices)
	dp := make([]int, 2)
	dp[0], dp[1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], dp[0]-prices[i])
	}
	return max(dp[0], dp[1])
}
