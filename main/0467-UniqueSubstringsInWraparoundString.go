// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/22 14:31
// -----------------------------------------------
package main

// 动态规划

func findSubstringInWraproundString(s string) int {
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := [26]int{}
	cnt := 0
	for i, ch := range s {
		if i > 0 && (byte(ch)-s[i-1]+26)%26 == 1 {
			cnt++
		} else {
			cnt = 1
		}
		dp[ch-'a'] = maxVal(dp[ch-'a'], cnt)
	}
	res := 0
	for i := 0; i < 26; i++ {
		res += dp[i]
	}
	return res
}
