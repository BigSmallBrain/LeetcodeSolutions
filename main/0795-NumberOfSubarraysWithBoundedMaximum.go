// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/19 22:29
// -----------------------------------------------
package main

// 滑动窗口 todo

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	last1, last2 := -1, -1
	res := 0
	for i, num := range nums {
		if num >= left && num <= right {
			last1 = i
		} else if num > right {
			last2 = i
			last1 = -1
		}
		if last1 != -1 {
			res += last1 - last2
		}
	}
	return res
}
