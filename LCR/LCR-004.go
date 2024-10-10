// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/10/10 13:43
// -----------------------------------------------
package main

func singleNumber(nums []int) int {
	count := make([]int32, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			count[i] += int32(num) >> i & int32(1)
		}
	}
	var res int32
	for i := 0; i < 32; i++ {
		res += count[i] % int32(3) << i
	}
	return int(res)
}
