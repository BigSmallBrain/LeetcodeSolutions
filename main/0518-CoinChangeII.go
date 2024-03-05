// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/4 21:09
// -----------------------------------------------
package main

// 动态规划 背包问题

func change(amount int, coins []int) (res int) {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i < amount+1; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
