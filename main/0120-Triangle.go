// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/22 10:36
// -----------------------------------------------
package main

import "math"

// 动态规划

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	dp[0] = make([]int, 1)
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}

	}
	res := math.MaxInt
	for i := 0; i < len(dp[m-1]); i++ {
		res = min(res, dp[m-1][i])
	}
	return res
}
