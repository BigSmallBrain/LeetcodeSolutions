// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/26 11:25
// -----------------------------------------------
package main

import "sort"

// 动态规划 排序

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	dp[0][0] = matrix[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] ^ matrix[0][i]
	}
	res := make([]int, 0)
	res = append(res, dp[0]...)
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = dp[i-1][0] ^ matrix[i][0]
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j-1] ^ dp[i-1][j] ^ dp[i][j-1] ^ matrix[i][j]
		}
		res = append(res, dp[i]...)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] > res[j]
	})
	return res[k-1]
}
