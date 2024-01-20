// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/20 19:39
// -----------------------------------------------
package main

// 动态规划

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	res := 0
	for i := 3; i < n+1; i++ {
		res = a + b + c
		a = b
		b = c
		c = res
	}
	return res
}
