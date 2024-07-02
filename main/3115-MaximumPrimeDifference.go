// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/2 17:41
// -----------------------------------------------
package main

// 数组

func maximumPrimeDifference(nums []int) int {
	isPrime := func(x int) bool {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return x >= 2
	}
	l := 0
	for !isPrime(nums[l]) {
		l++
	}
	r := len(nums) - 1
	for !isPrime(nums[r]) {
		r--
	}
	return r - l
}
