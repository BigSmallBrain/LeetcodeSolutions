// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/11 10:46
// -----------------------------------------------
package main

// 动态规划 背包问题

func numRollsToTarget(n int, k int, target int) int {
	mod := int(1e9 + 7)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < n+1; i++ {
		for j := target; j >= 0; j-- {
			dp[j] = 0
			for flag := 1; flag <= k; flag++ {
				if j-flag < 0 {
					break
				}
				dp[j] = (dp[j] + dp[j-flag]) % mod
			}
		}
	}
	return dp[target]
}
