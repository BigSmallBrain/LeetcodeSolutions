// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/10 19:41
// -----------------------------------------------
package main

// 动态规划

func robII(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := func(start, end int) int {
		a, b := nums[start], max(nums[start], nums[start+1])
		for i := start + 2; i <= end; i++ {
			temp := max(a+nums[i], b)
			a = b
			b = temp
		}
		return b
	}
	return max(dp(0, n-2), dp(1, n-1))
}
