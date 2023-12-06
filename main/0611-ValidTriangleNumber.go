// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/5 21:40
// -----------------------------------------------
package main

import "sort"

// 双指针 二分查找

func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	res := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			pos := sort.SearchInts(nums[j+1:], nums[i]+nums[j])
			res += pos
		}
	}
	return res
}
