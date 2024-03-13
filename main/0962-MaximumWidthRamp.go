// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/17 21:34
// -----------------------------------------------
package main

// 单调栈

func maxWidthRamp(nums []int) (res int) {
	n := len(nums)
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		if len(stack) == 0 || nums[i] > nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] <= nums[stack[len(stack)-1]] {
			res = max(res, stack[len(stack)-1]-i)
			stack = stack[:len(stack)-1]
		}
	}
	return
}
