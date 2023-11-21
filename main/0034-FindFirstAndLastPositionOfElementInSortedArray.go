// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/21 21:35
// -----------------------------------------------
package main

// 二分查找

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			l, r = mid, mid
			for l >= 0 && nums[l] == target {
				l--
			}
			for r < len(nums) && nums[r] == target {
				r++
			}
			return []int{l + 1, r - 1}
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return []int{-1, -1}
}
