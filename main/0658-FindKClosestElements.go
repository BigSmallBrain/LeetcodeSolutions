// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/1 20:55
// -----------------------------------------------
package main

import "sort"

// 二分查找 双指针

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)

	r := sort.SearchInts(arr, x)
	l := r - 1
	for ; k > 0; k-- {
		if l < 0 {
			r++
		} else if r >= n || x-arr[l] <= arr[r]-x {
			l--
		} else {
			r++
		}
	}

	return arr[l+1 : r]
}
