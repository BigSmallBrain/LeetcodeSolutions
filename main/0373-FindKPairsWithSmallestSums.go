// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/30 13:44
// -----------------------------------------------
package main

import "sort"

// 二分查找

func kSmallestPairs(nums1 []int, nums2 []int, k int) (res [][]int) {
	m, n := len(nums1), len(nums2)
	left, right := nums1[0]+nums2[0], nums1[m-1]+nums2[n-1]
	kthSum := left + sort.Search(right-left, func(sum int) bool { // 二分查找kthSum
		sum += left
		counter := 0
		i, j := 0, n-1
		for i < m && j >= 0 {
			if nums1[i]+nums2[j] > sum {
				j--
			} else {
				counter += j + 1
				i++
			}
		}
		return counter >= k
	})
	// 小于kthSum
	i := n - 1
	for _, num1 := range nums1 {
		for i >= 0 && num1+nums2[i] >= kthSum {
			i--
		}
		for _, num2 := range nums2[:i+1] {
			res = append(res, []int{num1, num2})
			if len(res) == k {
				return
			}
		}
	}
	// 等于kthSum
	i = n - 1
	for _, num1 := range nums1 {
		for i >= 0 && num1+nums2[i] > kthSum {
			i--
		}
		for j := i; j >= 0 && num1+nums2[j] == kthSum; j-- {
			res = append(res, []int{num1, nums2[j]})
			if len(res) == k {
				return
			}
		}
	}
	return
}
