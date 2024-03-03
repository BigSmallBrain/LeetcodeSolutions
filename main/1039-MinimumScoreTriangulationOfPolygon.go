// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/3 10:52
// -----------------------------------------------
package main

import "math"

// 动态规划

func minScoreTriangulation(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			minScore := math.MaxInt
			for k := i + 1; k < j; k++ {
				minScore = min(minScore, nums[i]*nums[k]*nums[j]+dp[i][k]+dp[k][j])
			}
			dp[i][j] = minScore
		}
	}
	return dp[0][n-1]
}
