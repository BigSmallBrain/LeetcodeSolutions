// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/24 21:43
// -----------------------------------------------
package main

// 动态规划

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, 2)
	for i := 2; i < n+1; i++ {
		dp[0], dp[1] = dp[1], min(dp[0]+cost[i-2], dp[1]+cost[i-1])
	}
	return dp[1]
}
