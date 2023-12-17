// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/17 19:19
// -----------------------------------------------
package main

// 滑动窗口

func longestOnes(nums []int, k int) int {
	n := len(nums)
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	l, r := 0, 0
	zeros := 0
	res := 0
	for r < n {
		if nums[r] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[l] == 0 {
				zeros--
			}
			l++
		}
		res = maxVal(res, r-l+1)
		r++
	}
	return res
}
