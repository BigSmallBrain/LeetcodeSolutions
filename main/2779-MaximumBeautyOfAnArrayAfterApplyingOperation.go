// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/16 16:12
// -----------------------------------------------
package main

import "sort"

// 双指针

func maximumBeauty(nums []int, k int) (res int) {
	n := len(nums)
	sort.Ints(nums)
	l, r := 0, 0
	for r < n {
		for r < n && nums[r]-nums[l] <= 2*k {
			r++
		}
		res = max(res, r-l)
		l++
	}
	return
}
