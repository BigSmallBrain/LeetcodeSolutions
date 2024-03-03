// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/1 20:45
// -----------------------------------------------
package main

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([]bool, n)
	start, end := 0, 0
	for i := n - 1; i >= 0; i-- {
		dp[i] = true
		pre := false
		for j := i + 1; j < n; j++ {
			if s[i] != s[j] {
				pre, dp[j] = false, false
			} else {
				if j-i < 3 {
					pre, dp[j] = true, true
				} else {
					pre, dp[j] = dp[j], pre
				}
			}
			if dp[j] && j-i > end-start {
				start, end = i, j
			}
		}
	}
	return s[start : end+1]
}
