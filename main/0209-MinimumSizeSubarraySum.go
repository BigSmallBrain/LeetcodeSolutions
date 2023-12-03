// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/1 21:53
// -----------------------------------------------
package main

import (
	"math"
	"sort"
)

// 二分查找 前缀和

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)

	minVal := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for i := 1; i < n+1; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		tgt := sum[i] + target
		index := sort.SearchInts(sum, tgt)
		if index < n+1 {
			res = minVal(res, index-i)
		}
	}

	if res > n {
		return 0
	}
	return res
}
