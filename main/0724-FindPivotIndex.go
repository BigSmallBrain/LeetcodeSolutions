// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/8 16:17
// -----------------------------------------------
package main

func pivotIndex(nums []int) int {
	n := len(nums)

	leftSum, rightSum := 0, 0
	for i := 1; i < n; i++ {
		rightSum += nums[i]
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			leftSum += nums[i-1]
			rightSum -= nums[i]
		}
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
