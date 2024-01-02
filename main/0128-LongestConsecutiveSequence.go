// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/2 19:31
// -----------------------------------------------
package main

import "sort"

// 排序

func longestConsecutive(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	sort.Ints(nums)
	cnt := 1
	res := 1
	for i := 0; i < n-1; i++ {
		if nums[i+1] == nums[i] {
			continue
		}
		if nums[i+1] != nums[i]+1 {
			cnt = 1
		} else {
			cnt++
			if cnt > res {
				res = cnt
			}
		}
	}
	return res
}
