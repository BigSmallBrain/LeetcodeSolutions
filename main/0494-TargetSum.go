// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/20 20:32
// -----------------------------------------------
package main

// 动态规划 背包问题

func findTargetSumWays(nums []int, target int) (res int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return
	}
	diff >>= 1
	dp := make([]int, diff+1)
	dp[0] = 1
	for _, num := range nums {
		for i := diff; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[diff]
}
