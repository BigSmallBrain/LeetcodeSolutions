// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/22 20:50
// -----------------------------------------------
package main

import "sort"

// 暴力解法
func mergeNums(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)
}
