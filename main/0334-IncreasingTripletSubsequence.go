// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/7 16:48
// -----------------------------------------------
package main

import "math"

// 贪心算法

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	first, second := nums[0], math.MaxInt
	for i := 1; i < n; i++ {
		if nums[i] > second {
			return true
		} else if nums[i] > first {
			second = nums[i]
		} else {
			first = nums[i]
		}
	}
	return false
}
