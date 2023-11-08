// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/8 15:28
// -----------------------------------------------
package main

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}

	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k+1:])
}

func reverse(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
