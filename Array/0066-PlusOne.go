// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/22 15:56
// -----------------------------------------------
package main

func plusOne(digits []int) []int {
	flag := false
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i >= 0; i-- {
		if flag {
			digits[i] += 1
		}
		if digits[i] > 9 {
			digits[i] %= 10
			flag = true
		} else {
			flag = false
		}

		if flag && i == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
