// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/24 10:51
// -----------------------------------------------
package main

// 贪心 单调栈

func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0)
	n := len(nums)
	for i, num := range nums {
		for len(stack) > 0 && n-(i+1) >= k-len(stack) && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, num)
	}
	return stack[:k]
}
