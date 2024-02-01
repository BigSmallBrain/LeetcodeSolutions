// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/1 17:49
// -----------------------------------------------
package main

// 原地交换

func findDisappearedNumbers(nums []int) (res []int) {
	for _, num := range nums {
		for num != nums[num-1] {
			num, nums[num-1] = nums[num-1], num
		}
	}
	for i, num := range nums {
		if num != i+1 {
			res = append(res, i+1)
		}
	}
	return
}
