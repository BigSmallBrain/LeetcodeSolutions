// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/5 21:11
// -----------------------------------------------
package main

// 双指针

func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1

	res := make([]int, 0)

	for l <= r {
		left := nums[l] * nums[l]
		right := nums[r] * nums[r]
		if left > right {
			res = append(res, left)
			l++
		} else {
			res = append(res, right)
			r--
		}
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
