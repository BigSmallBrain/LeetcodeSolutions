// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/18 22:06
// -----------------------------------------------
package main

// 动态规划 前缀和

func canReachII(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	dp, pre := make([]int, n), make([]int, n)
	dp[0] = 1
	for i := 0; i < minJump; i++ {
		pre[i] = 1
	}
	for i := minJump; i < n; i++ {
		left, right := i-maxJump, i-minJump
		if s[i] == '0' {
			sum := pre[right]
			if left > 0 {
				sum -= pre[left-1]
			}
			if sum > 0 {
				dp[i] = 1
			}
		}
		pre[i] = pre[i-1] + dp[i]
	}
	return dp[n-1] > 0
}
