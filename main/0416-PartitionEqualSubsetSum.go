// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/22 18:43
// -----------------------------------------------
package main

// 动态规划 背包问题

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = max(dp[i], dp[i-num]+num)
		}
	}
	return dp[target] == target
}
