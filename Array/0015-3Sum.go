// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/17 10:02
// -----------------------------------------------
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	res := make([][]int, 0)
	resMap := make(map[string][]int)
	// 遍历
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for right > left {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				tempRes := []int{nums[i], nums[left], nums[right]}
				tempResStr := fmt.Sprintf("%v", tempRes)
				resMap[tempResStr] = tempRes
				right--
				left++
			}
			// 结果偏小
			if sum < 0 {
				left++
			}
			// 结果偏大
			if sum > 0 {
				right--
			}
		}

	}
	// 转换结果
	for _, value := range resMap {
		res = append(res, value)
	}
	return res
}
