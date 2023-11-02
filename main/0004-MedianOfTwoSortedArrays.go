// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/23 10:30
// -----------------------------------------------
package main

import "sort"

// 划分数组
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)

	left, right := 0, len(nums)-1
	for left < right {
		left++
		right--
	}

	if left == right {
		return float64(nums[left])
	} else {
		return float64(nums[left]+nums[right]) / 2
	}
}
