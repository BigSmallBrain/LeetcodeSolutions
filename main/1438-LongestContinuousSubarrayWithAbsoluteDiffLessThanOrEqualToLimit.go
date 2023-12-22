// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/22 15:40
// -----------------------------------------------
package main

// 滑动窗口 单调队列 todo

func longestSubarrayWithAbsoluteDiff(nums []int, limit int) int {
	n := len(nums)
	l, r := 0, 0
	queMax, queMin := make([]int, 0), make([]int, 0)
	res := 0
	for r < n {
		for len(queMin) > 0 && queMin[len(queMin)-1] > nums[r] {
			queMin = queMin[:len(queMin)-1]
		}
		queMin = append(queMin, nums[r])
		for len(queMax) > 0 && queMax[len(queMax)-1] < nums[r] {
			queMax = queMax[:len(queMax)-1]
		}
		queMax = append(queMax, nums[r])
		for len(queMin) > 0 && len(queMax) > 0 && queMax[0]-queMin[0] > limit {
			if nums[l] == queMin[0] {
				queMin = queMin[1:]
			}
			if nums[l] == queMax[0] {
				queMax = queMax[1:]
			}
			l++
		}
		if res < r-l+1 {
			res = r - l + 1
		}
		r++
	}
	return res
}
