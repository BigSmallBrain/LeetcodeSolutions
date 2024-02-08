// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/8 17:37
// -----------------------------------------------
package main

import "sort"

// 动态规划 二分查找 贪心

func lengthOfLIS(nums []int) int {
	temp := make([]int, 0)
	for _, num := range nums {
		i := sort.SearchInts(temp, num)
		if i == len(temp) {
			temp = append(temp, num)
		} else {
			temp[i] = num
		}
	}
	return len(temp)
}
