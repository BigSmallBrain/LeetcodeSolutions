// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/8 18:32
// -----------------------------------------------
package main

func findMaxConsecutiveOnes(nums []int) int {
	res := -1
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			flag += 1
			if flag > res {
				res = flag
			}
		} else {
			flag = 0
		}
	}
	return res
}
