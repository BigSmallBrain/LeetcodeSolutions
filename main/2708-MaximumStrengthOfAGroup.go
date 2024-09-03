// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/9/3 13:33
// -----------------------------------------------
package main

import "sort"

// 分类讨论

func maxStrength(nums []int) int64 {
	n := len(nums)
	if n == 1 && nums[0] < 0 { // 特殊情况
		return int64(nums[0])
	}
	sort.Ints(nums)
	res := int64(1)

	if nums[n-1] <= 0 {
		i := 0
		for ; i < n && nums[i] < 0; i += 2 { // 负数
			if i+1 < n && nums[i+1] < 0 {
				res *= int64(nums[i]) * int64(nums[i+1])
			} else {
				break
			}
		}
		if i == 0 {
			return 0
		}
		return res
	}

	for i := n - 1; i >= 0 && nums[i] > 0; i-- { // 正数
		res *= int64(nums[i])
	}
	for i := 0; i < n && nums[i] < 0; i += 2 { // 负数
		if i+1 < n && nums[i+1] < 0 {
			res *= int64(nums[i]) * int64(nums[i+1])
		} else {
			break
		}
	}
	return res
}
