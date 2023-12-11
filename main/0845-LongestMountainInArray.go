// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/7 16:19
// -----------------------------------------------
package main

// 双指针

func longestMountain(arr []int) int {
	n := len(arr)

	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 3
	for i := 1; i < n-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			l, r := i, i
			for l > 0 && arr[l-1] < arr[l] {
				l--
			}
			for r < n-1 && arr[r+1] < arr[r] {
				r++
			}
			res = maxVal(res, r-l+1)
		}
	}
	return res
}
