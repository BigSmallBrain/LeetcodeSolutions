// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/6 11:05
// -----------------------------------------------
package main

// 动态规划 背包问题

func numSquares(n int) int {
	squares := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = n + 1
	}
	dp[0] = 0
	for _, square := range squares {
		for i := square; i < n+1; i++ {
			dp[i] = min(dp[i], dp[i-square]+1)
		}
	}
	return dp[n]
}
