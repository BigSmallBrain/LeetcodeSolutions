// Package Stack
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/2 18:02
// -----------------------------------------------
package main

// 单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	tempRes := make(map[int]int, len(nums2))
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] >= nums2[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			tempRes[nums2[i]] = -1
		} else {
			tempRes[nums2[i]] = nums2[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}

	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = tempRes[nums1[i]]
	}
	return res
}
