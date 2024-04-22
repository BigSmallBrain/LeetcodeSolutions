// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/22 10:27
// -----------------------------------------------
package main

import "math"

// 动态规划

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[n-1]
}
