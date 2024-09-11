// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/9/11 23:09
// -----------------------------------------------
package main

import "sort"

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Ints(nums)
	res, length := 1, 1
	last := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == last+1 {
			length++
			res = max(res, length)
		} else {
			length = 1
		}
		last = nums[i]
	}
	return res
}
