// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/10 21:48
// -----------------------------------------------
package main

// 位运算

func findErrorNums(nums []int) []int {
	var xor int
	for i, num := range nums {
		xor = xor ^ num ^ (i + 1)
	}
	flag := xor & (-xor)
	var a, b int
	for i, num := range nums {
		if flag&num == 0 {
			a ^= num
		}
		if flag&(i+1) == 0 {
			a ^= i + 1
		}
	}
	b = a ^ xor
	for _, num := range nums {
		if a == num {
			return []int{a, b}
		}
	}
	return []int{b, a}
}
