// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/4 18:43
// -----------------------------------------------
package main

// 位运算

func reverseBits(num uint32) (res uint32) {
	for i := 0; i < 32 && num > 0; i++ {
		res |= num & uint32(1) << (32 - i)
		num >>= 1
	}
	return
}
