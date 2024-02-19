// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/18 19:04
// -----------------------------------------------
package main

// 贪心算法

func jump(nums []int) int {
	n := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < n-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
