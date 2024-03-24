// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/24 22:13
// -----------------------------------------------
package main

func differenceOfSum(nums []int) (res int) {
	for _, num := range nums {
		flag := 10
		temp := num
		for temp >= 10 {
			temp /= 10
			digit := temp % 10
			res += digit*flag - digit
			flag *= 10
		}
	}
	return
}
