// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/29 20:23
// -----------------------------------------------
package main

// 哈希

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	leftSum := map[int]int{}
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			leftSum[v1+v2]++
		}
	}
	cnt := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			cnt += leftSum[-v3-v4]
		}
	}
	return cnt
}
