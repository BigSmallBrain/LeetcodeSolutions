// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/23 10:59
// -----------------------------------------------
package main

// 数学

func trailingZeroes(n int) (res int) {
	for n > 0 {
		n /= 5
		res += n
	}
	return
}
