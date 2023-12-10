// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/10 17:29
// -----------------------------------------------
package main

// 滑动窗口

func findMaxAverage(nums []int, k int) float64 {
	if k > len(nums) {
		return 0
	}

	l, r := 0, k-1
	sum := 0
	for i := l; i <= r; i++ {
		sum += nums[i]
	}
	res := sum

	for l, r = l+1, r+1; r < len(nums); l, r = l+1, r+1 {
		sum = sum - nums[l-1] + nums[r]
		if sum > res {
			res = sum
		}
	}
	return float64(res) / float64(k)
}
