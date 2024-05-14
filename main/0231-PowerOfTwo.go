// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/21 22:04
// -----------------------------------------------
package main

// 位运算

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1) == 0)
}
