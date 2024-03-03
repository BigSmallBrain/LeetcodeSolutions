// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/1 19:45
// -----------------------------------------------
package main

// 动态规划

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		pre := 0
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				pre, dp[j] = dp[j], pre+2
			} else {
				pre = dp[j]
				dp[j] = max(dp[j], dp[j-1])
			}
		}
	}
	return dp[n-1]
}
