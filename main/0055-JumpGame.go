// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/15 19:01
// -----------------------------------------------
package main

// 贪心

func canJump(nums []int) bool {
	jump := nums[0]
	for i := 1; i <= jump && i < len(nums); i++ {
		jump = max(jump, i+nums[i])
	}
	return jump >= len(nums)-1
}
