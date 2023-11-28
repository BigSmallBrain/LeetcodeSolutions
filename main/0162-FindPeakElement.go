// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/24 20:45
// -----------------------------------------------
package main

import "math"

// 二分查找

func findPeakElement(nums []int) int {
	n := len(nums)
	get := func(i int) int {
		if i == -1 || i == n-1 {
			return math.MinInt
		}
		return nums[i]
	}
	l, r := 0, n-1
	for {
		mid := (l + r) >> 1
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}
