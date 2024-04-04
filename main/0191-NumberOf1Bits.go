// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/3 20:23
// -----------------------------------------------
package main

import "math/bits"

// 位运算

func hammingWeight(n int) int {
	return bits.OnesCount(uint(n))
}
