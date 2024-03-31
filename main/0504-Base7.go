// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/31 20:48
// -----------------------------------------------
package main

import "slices"

// 位运算

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	nums := make([]byte, 0)
	flag := false
	if num < 0 {
		num = -num
		flag = true
	}
	for num > 0 {
		nums = append(nums, byte('0'+num%7))
		num /= 7
	}
	if flag {
		nums = append(nums, '-')
	}
	slices.Reverse(nums)
	return string(nums)
}
