// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/15 20:45
// -----------------------------------------------
package main

// 滑动窗口 前缀和

func minOperations(nums []int, x int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum < x {
		return -1
	}

	leftSum, rightSum := 0, sum
	res := n + 1
	r := 0
	for l := -1; l < n; l++ {
		if l > -1 {
			leftSum += nums[l]
		}
		for r < n && leftSum+rightSum > x {
			rightSum -= nums[r]
			r++
		}
		if leftSum+rightSum == x && (l+1)+(n-r) < res {
			res = (l + 1) + (n - r)
		}
	}
	if res > n {
		return -1
	}
	return res
}
