// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/8 10:05
// -----------------------------------------------
package main

import "sort"

// 二分查找

func numSubseq(nums []int, target int) (res int) {
	n := len(nums)
	mod := int(1e9 + 7)
	pow := make([]int, n)
	pow[0] = 1
	for i := 1; i < n; i++ {
		pow[i] = (pow[i-1] * 2) % mod
	}
	sort.Ints(nums)
	for i, num := range nums {
		if num*2 > target {
			continue
		}
		maxVal := target - num
		index := sort.SearchInts(nums, maxVal)
		for index < n && nums[index] == maxVal { // 寻找最右侧大于maxVal的索引
			index++
		}
		if index-1 < i {
			continue
		}
		res = (res + pow[index-i-1]) % mod
	}
	return
}
