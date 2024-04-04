// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/4 19:34
// -----------------------------------------------
package main

// 位运算

func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	flag := 1
	for i := num; i > 0; i >>= 1 {
		flag <<= 1
	}
	return num ^ (flag - 1)
}
