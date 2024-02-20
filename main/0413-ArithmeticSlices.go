// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/19 18:06
// -----------------------------------------------
package main

// 动态规划

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := 0
	flag := 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp += i - flag - 1
		} else {
			flag = i - 1
		}
	}
	return dp
}
