// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/17 22:09
// -----------------------------------------------
package main

// 滑动窗口

func maximumUniqueSubarray(nums []int) int {
	n := len(nums)
	numMap := map[int]struct{}{}
	l, r := 0, 0
	res := 0
	sum := 0
	for r < n {
		for _, ok := numMap[nums[r]]; l <= r && ok; _, ok = numMap[nums[r]] {
			delete(numMap, nums[l])
			sum -= nums[l]
			l++
		}
		numMap[nums[r]] = struct{}{}
		sum += nums[r]
		if sum > res {
			res = sum
		}
		r++
	}
	return res
}
