// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/31 16:44
// -----------------------------------------------
package main

import (
	"math"
)

// TODO 单调栈
func sumSubarrayMins(arr []int) int {
	N := int(math.Pow(10, 9)) + 7
	n := len(arr)
	left, right := make([]int, n), make([]int, n)

	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[i] <= arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[i] < arr[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := 0
	for i := 0; i < n; i++ {
		res += arr[i] * (i - left[i]) * (right[i] - i)
		res %= N
	}
	return res
}
