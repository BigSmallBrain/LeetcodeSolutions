// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/28 21:06
// -----------------------------------------------
package main

// 交换位置

func sortArrayByParityII(nums []int) []int {
	even := 0
	for odd := 1; odd < len(nums); odd += 2 {
		if nums[odd]%2 != 1 {
			for nums[even]%2 == 0 {
				even += 2
			}
			nums[odd], nums[even] = nums[even], nums[odd]
		}
	}
	return nums
}
