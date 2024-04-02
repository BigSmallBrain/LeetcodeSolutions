// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/2 9:45
// -----------------------------------------------
package main

// 位运算

func singleNumberIII(nums []int) []int {
	if len(nums) == 2 {
		return nums
	}
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	res := make([]int, 2)
	flag := xor & (-xor)
	for _, num := range nums {
		if num&flag == 0 {
			res[0] ^= num
		}
	}
	res[1] = xor ^ res[0]
	return res
}
