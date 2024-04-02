// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/2 10:01
// -----------------------------------------------
package main

// 位运算

func missingNumber(nums []int) (res int) {
	for i, num := range nums {
		res = res ^ i ^ num
	}
	return res ^ len(nums)
}
