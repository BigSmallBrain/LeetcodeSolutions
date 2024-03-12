// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/12 15:01
// -----------------------------------------------
package main

// 动态规划

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	words := map[string]struct{}{}
	lengths := make([]bool, 21)
	for _, w := range wordDict {
		words[w] = struct{}{}
		lengths[len(w)] = true
	}
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i < n+1; i++ {
		for length, flag := range lengths {
			if flag && i-length >= 0 {
				_, ok := words[s[i-length:i]]
				dp[i] = dp[i] || (dp[i-length] && ok)
			}
		}
	}
	return dp[n]
}
