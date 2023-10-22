// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/16 10:02
// -----------------------------------------------
package main

func twoSum(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	numsMap := make(map[int]int)

	for i, num := range nums {
		temp := target - num

		if index, ok := numsMap[temp]; ok {
			return []int{index, i}
		}

		numsMap[num] = i
	}

	return []int{}
}
