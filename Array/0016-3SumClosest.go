// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/18 9:35
// -----------------------------------------------
package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var res int
	tempRes := math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for right > left {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return sum
			}

			if sum < target {
				left++
				temp := int(math.Abs(float64(sum - target)))
				if tempRes >= temp {
					tempRes = temp
					res = sum
				}
			}

			if sum > target {
				right--
				temp := int(math.Abs(float64(sum - target)))
				if tempRes >= temp {
					tempRes = temp
					res = sum
				}
			}
		}
	}
	return res
}
