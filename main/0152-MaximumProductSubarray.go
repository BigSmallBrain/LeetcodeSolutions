// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/13 17:47
// -----------------------------------------------
package main

// 动态规划

func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mn*nums[i], mx*nums[i], nums[i])
		minF = min(mn*nums[i], mx*nums[i], nums[i])
		ans = max(maxF, ans)
	}
	return ans
}
