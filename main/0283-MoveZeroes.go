// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/11 22:08
// -----------------------------------------------
package main

func moveZeroes(nums []int) {
	n := len(nums)
	l := 0

	for r := 0; r < n; r++ {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
	}

	for ; l < n; l++ {
		nums[l] = 0
	}
}
