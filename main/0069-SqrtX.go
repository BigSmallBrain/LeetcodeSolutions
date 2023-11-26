// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/25 20:48
// -----------------------------------------------
package main

func mySqrt(x int) int {
	l, r := 0, x
	res := -1
	for l <= r {
		mid := (l + r) >> 1
		temp := mid * mid
		if temp == x {
			return mid
		}
		if temp < x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
