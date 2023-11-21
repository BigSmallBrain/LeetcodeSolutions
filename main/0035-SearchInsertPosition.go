// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/21 21:21
// -----------------------------------------------
package main

import "sort"

// 二分查找

func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(x int) bool {
		return nums[x] >= target
	})
}
