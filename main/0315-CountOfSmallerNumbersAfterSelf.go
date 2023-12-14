// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/14 19:45
// -----------------------------------------------
package main

// TODO

func countSmaller(nums []int) []int {
	minVal, maxVal, n := nums[0], nums[0], len(nums)
	for i := 1; i < n; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}
	length := maxVal - minVal + 2
	has, res := make([]int, length), make([]int, n)
	for i := n - 1; i >= 0; i-- {
		nv := nums[i] - minVal
		tep := 0
		for nv > 0 {
			if has[nv] > 0 {
				tep += has[nv]
			}
			nv &= nv - 1
		}
		nv = nums[i] - minVal + 1
		for nv < length {
			has[nv]++
			nv += nv & (-nv)
		}
		res[i] = tep
	}
	return res
}
