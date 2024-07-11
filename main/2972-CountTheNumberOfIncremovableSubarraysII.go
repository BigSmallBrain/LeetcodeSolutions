// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/11 17:19
// -----------------------------------------------
package main

// 双指针

func incremovableSubarrayCount(nums []int) int64 {
	n := len(nums)
	var res int64 = 0
	l := 0
	for l < n {
		if l > 0 && nums[l] <= nums[l-1] {
			break
		}
		l++
	}
	if l == n { // 特殊情况
		return int64(n) * int64(n+1) / 2
	}
	res += int64(l + 1)
	for r := n - 1; r > 0; r-- {
		if r < n-1 && nums[r] >= nums[r+1] {
			break
		}
		for l > 0 && nums[l-1] >= nums[r] {
			l--
		}
		res += int64(l + 1)
	}
	return res
}
