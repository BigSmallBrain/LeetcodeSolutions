// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/5 21:38
// -----------------------------------------------
package main

// 动态规划

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	idx2, idx3, idx5 := 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(dp[idx2]*2, dp[idx3]*3, dp[idx5]*5)
		if dp[i] == dp[idx2]*2 {
			idx2++
		}
		if dp[i] == dp[idx3]*3 {
			idx3++
		}
		if dp[i] == dp[idx5]*5 {
			idx5++
		}
	}
	return dp[n-1]
}
