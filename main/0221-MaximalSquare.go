// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/25 10:23
// -----------------------------------------------
package main

// 动态规划

func maximalSquare(matrix [][]byte) (res int) {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				temp[j] = 1
				res = 1
			}
		}
		dp[i] = temp
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				res = max(res, dp[i][j])
			}
		}

	}
	return res * res
}
