// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/19 20:43
// -----------------------------------------------
package main

// 动态规划

func climbStairs(n int) (res int) {
	if n < 3 {
		return n
	}
	dp1, dp2 := 2, 1
	for i := 3; i < n+1; i++ {
		res = dp1 + dp2
		dp2 = dp1
		dp1 = res
	}
	return
}
