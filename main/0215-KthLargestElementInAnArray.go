// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/12 11:23
// -----------------------------------------------
package main

import "math/rand"

// 快速排序
func findKthLargest(nums []int, k int) int {
	quickSelect(nums)
	return nums[len(nums)-k]
}

func quickSelect(nums []int) {
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

	quickSelect(nums[:j+1])
	quickSelect(nums[j+1:])
}
