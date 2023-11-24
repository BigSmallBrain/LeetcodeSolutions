// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/24 19:20
// -----------------------------------------------
package main

import "sort"

// 二分查找 划分数组

func searchInRotatedSortedArray(nums []int, target int) int {
	index := sort.Search(len(nums), func(i int) bool {
		if nums[i] >= nums[0] {
			return target >= nums[0] && nums[i] >= target
		} else {
			return target >= nums[0] || nums[i] >= target
		}
	})
	if index < len(nums) && nums[index] == target {
		return index
	}
	return -1
}
