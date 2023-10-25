// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/25 19:37
// -----------------------------------------------
package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	length := len(nums)
	dp := make([]int, length)

	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < length; i++ {
		dp[i] = int(math.Max(float64(dp[i-1]+nums[i]), float64(nums[i])))
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}
