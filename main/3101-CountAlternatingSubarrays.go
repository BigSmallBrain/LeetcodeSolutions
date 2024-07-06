// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/6 11:55
// -----------------------------------------------
package main

// 动态规划

func countAlternatingSubarrays(nums []int) (res int64) {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	res += int64(dp[0])
	for i := 1; i < n; i++ {
		dp[i] = 1
		if nums[i] != nums[i-1] {
			dp[i] += dp[i-1]
		}
		res += int64(dp[i])
	}
	return
}
