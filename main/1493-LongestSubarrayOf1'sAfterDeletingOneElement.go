// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/18 16:34
// -----------------------------------------------
package main

// 滑动窗口

func longestSubarray(nums []int) int {
	n := len(nums)
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res := 0
	deleteFlag := 0
	l, r := 0, 0
	for r < n {
		if nums[r] == 0 {
			deleteFlag++
		}
		for deleteFlag > 1 {
			if nums[l] == 0 {
				deleteFlag--
			}
			l++
		}
		res = maxVal(res, r-l-deleteFlag+1)
		r++
	}
	if deleteFlag == 0 {
		return res - 1
	}
	return res
}
