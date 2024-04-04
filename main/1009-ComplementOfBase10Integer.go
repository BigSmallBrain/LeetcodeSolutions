// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/4 19:14
// -----------------------------------------------
package main

// 位运算

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	flag := 1
	for i := n; i > 0; i >>= 1 {
		flag <<= 1
	}
	return n ^ (flag - 1)
}
