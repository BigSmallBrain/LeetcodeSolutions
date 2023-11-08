// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/8 18:37
// -----------------------------------------------
package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	rightRes := make([]int, n)

	left, right := 1, 1
	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			right *= nums[i+1]
		}
		rightRes[i] = right
	}

	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 {
			left *= nums[i-1]
		}
		res = append(res, left*rightRes[i])
	}

	return res
}
