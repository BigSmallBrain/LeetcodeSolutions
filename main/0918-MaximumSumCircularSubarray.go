// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/17 13:34
// -----------------------------------------------
package main

// 动态规划

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	dp := make([]int, 2)
	// max min
	dp[0], dp[1] = nums[0], nums[0]
	maxSum, minSum := nums[0], nums[0]
	totalSum := nums[0]
	for i := 1; i < n; i++ {
		dp[0] = max(dp[0]+nums[i], nums[i])
		dp[1] = min(dp[1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[0])
		minSum = min(minSum, dp[1])
		totalSum += nums[i]
	}
	if maxSum < 0 {
		return maxSum
	}
	return max(maxSum, totalSum-minSum)
}
