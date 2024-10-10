// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/10/10 13:37
// -----------------------------------------------
package main

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i < n+1; i++ {
		if i&1 == 0 {
			dp[i] = dp[i>>1]
		} else {
			dp[i] = dp[(i-1)>>1] + 1
		}
	}
	return dp
}
