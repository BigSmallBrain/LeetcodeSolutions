// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/1 12:49
// -----------------------------------------------
package main

// 位运算

func singleNumberII(nums []int) int {
	counter := [32]int32{}
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			counter[i] += int32(num) >> i & int32(1)
		}
	}
	var res int32
	for i := 0; i < 32; i++ {
		res += counter[i] % int32(3) << i
	}
	return int(res)
}
