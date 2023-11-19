// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/16 19:52
// -----------------------------------------------
package main

import (
	"math"
)

// 桶排序

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	minVal, maxVal := math.MaxInt, -1
	for i := 0; i < n; i++ {
		if minVal > nums[i] {
			minVal = nums[i]
		}
		if maxVal < nums[i] {
			maxVal = nums[i]
		}
	}

	size := (maxVal-minVal)/n + 1
	cnt := (maxVal-minVal)/size + 1

	bucket := make([][]int, cnt)
	for i := 0; i < cnt; i++ {
		bucket[i] = append(bucket[i], []int{math.MaxInt, -1}...)
	}

	// 装入桶
	for i := 0; i < n; i++ {
		idx := (nums[i] - minVal) / size
		if nums[i] < bucket[idx][0] {
			bucket[idx][0] = nums[i]
		}
		if nums[i] > bucket[idx][1] {
			bucket[idx][1] = nums[i]
		}
	}

	res := -1
	for i := 1; i < len(bucket); i++ {
		if bucket[i][0] == math.MaxInt && bucket[i][1] == -1 {
			bucket[i][1] = bucket[i-1][1]
			continue
		}
		temp := bucket[i][0] - bucket[i-1][1]
		if res < temp {
			res = temp
		}
	}
	return res
}
