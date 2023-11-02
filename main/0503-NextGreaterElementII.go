// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/2 19:45
// -----------------------------------------------
package main

// 单调栈
func nextGreaterElements(nums []int) []int {
	flag := len(nums)
	nums = append(nums, nums[:len(nums)-1]...)
	n := len(nums)

	res := make([]int, n)
	stack := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = nums[stack[len(stack)-1]]
		}

		stack = append(stack, i)
	}

	return res[:flag]
}
