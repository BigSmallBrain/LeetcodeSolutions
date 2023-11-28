// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/27 21:26
// -----------------------------------------------
package main

import "math"

// 二分查找

func findNthDigit(n int) int {
	digit := 1
	for count := 9; n > digit*count; count *= 10 {
		n -= digit * count
		digit++
	}
	index := n - 1
	start := int(math.Pow10(digit - 1))

	num := start + index/digit
	digitIndex := index % digit
	return num / int(math.Pow10(digit-digitIndex-1)) % 10
}
