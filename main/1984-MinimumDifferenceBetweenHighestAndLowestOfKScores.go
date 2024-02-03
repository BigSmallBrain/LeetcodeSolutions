// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/3 23:28
// -----------------------------------------------
package main

import (
	"math"
	"sort"
)

// 滑动窗口

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	res := math.MaxInt
	for start, end := 0, k-1; end < len(nums); start, end = start+1, end+1 {
		if res > nums[end]-nums[start] {
			res = nums[end] - nums[start]
		}
	}
	return res
}
