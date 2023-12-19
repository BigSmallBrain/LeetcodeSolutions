// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/19 21:38
// -----------------------------------------------
package main

// 滑动窗口

func numSubarrayProductLessThanK(nums []int, k int) int {
	n := len(nums)

	res := 0
	l, r := 0, 0
	product := 1
	for r < n {
		product *= nums[r]
		for l <= r && product >= k {
			product /= nums[l]
			l++
		}
		res += r - l + 1
		r++
	}
	return res
}
