// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/10/11 14:45
// -----------------------------------------------
package main

import "sort"

func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i, num := range nums {
		index := sort.SearchInts(nums[i+1:], target-num)
		if i+index+1 < n && nums[i+index+1] == target-num {
			return []int{i, i + index + 1}
		}
	}
	return make([]int, 0)
}
