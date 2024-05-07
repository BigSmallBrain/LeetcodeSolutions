// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/7 16:15
// -----------------------------------------------
package main

// 动态规划

func divisorGame(n int) bool {
	if n == 1 {
		return false
	} else if n == 2 {
		return true
	}
	dp := make([]bool, n+1)
	dp[1], dp[2] = false, true
	for i := 3; i < n+1; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
