// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/18 18:19
// -----------------------------------------------
package main

// 动态规划

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, 2)
	dp[0], dp[1] = 1, 1
	front := s[0]
	for i := 1; i < n; i++ {
		temp := 0
		if s[i]-'0' >= 1 && s[i]-'0' <= 9 {
			temp += dp[1]
		}
		if (front-'0')*10+s[i]-'0' >= 10 && (front-'0')*10+s[i]-'0' <= 26 {
			temp += dp[0]
		}
		front = s[i]
		dp[0], dp[1] = dp[1], temp
	}
	return dp[1]
}
