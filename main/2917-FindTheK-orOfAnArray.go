// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/29 22:35
// -----------------------------------------------
package main

// 位运算

func findKOr(nums []int, k int) (res int) {
	counter := make([]int, 32)
	for i := 0; i < 32; i++ {
		for _, num := range nums {
			if num&(1<<i) > 0 {
				counter[i]++
			}
		}
	}
	for i := 0; i < 32; i++ {
		if counter[i] >= k {
			res += 1 << i
		}
	}
	return
}
