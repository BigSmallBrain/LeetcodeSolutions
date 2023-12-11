// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/11 20:13
// -----------------------------------------------
package main

// 双指针

func findLengthOfLCIS(nums []int) int {
	n := len(nums)

	res := 1
	l, r := 0, 0
	for r < n-1 {
		if nums[r+1] > nums[r] {
			r++
		} else {
			r++
			l = r
		}
		if res < r-l+1 {
			res = r - l + 1
		}
	}
	return res
}
