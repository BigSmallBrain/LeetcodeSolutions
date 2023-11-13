// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/13 18:43
// -----------------------------------------------
package main

import (
	"math/rand"
)

func majorityElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	sortNums(nums)
	return nums[n/2]
}

func sortNums(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	pivot := nums[rand.Intn(n)]

	i, j := -1, n
	for i < j {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 左部
	sortNums(nums[:j+1])
	// 右部
	sortNums(nums[j+1:])
}
