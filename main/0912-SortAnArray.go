// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/12 21:17
// -----------------------------------------------
package main

import "math/rand"

// 快速排序
func sortArray(nums []int) []int {
	quickSort(nums)
	return nums
}

func quickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 随机基准值
	pivotIndex := rand.Intn(n)
	nums[0], nums[pivotIndex] = nums[pivotIndex], nums[0]

	pivot := nums[0]
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
	quickSort(nums[:j+1])
	quickSort(nums[j+1:])
}
