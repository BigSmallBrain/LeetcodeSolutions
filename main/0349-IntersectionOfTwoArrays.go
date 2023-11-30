// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/30 20:41
// -----------------------------------------------
package main

import "sort"

// 排序 二分查找

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	n1, n2 := len(nums1), len(nums2)

	res := make([]int, 0)
	for i := 0; i < n1; i++ {
		if i > 0 && nums1[i] == nums1[i-1] {
			continue
		}
		index := sort.Search(n2-1, func(x int) bool {
			return nums2[x] >= nums1[i]
		})
		if index < n2 && nums2[index] == nums1[i] {
			res = append(res, nums1[i])
		}
	}
	return res
}
