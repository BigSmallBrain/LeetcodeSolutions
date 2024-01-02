// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/2 19:08
// -----------------------------------------------
package main

// 哈希表

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		} else {
			res = append(res, num)
		}
	}
	return res
}
