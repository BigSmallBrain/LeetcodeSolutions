// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/21 22:07
// -----------------------------------------------
package main

import "sort"

// 二分查找

func twoSumII(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		temp := numbers[i+1:]
		index, ok := sort.Find(len(temp), func(x int) int {
			return target - temp[x] - numbers[i]
		})
		if ok {
			return []int{i + 1, i + index + 2}

		}
	}
	return []int{-1, -1}
}
