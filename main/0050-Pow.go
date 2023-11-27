// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/27 20:52
// -----------------------------------------------
package main

// 快速幂

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / quickMul(x, -n)
	}
	return quickMul(x, n)
}

func quickMul(x float64, N int) float64 {
	res := 1.0

	contribute := x
	for N > 0 {
		if N%2 == 1 {
			res *= contribute
		}
		contribute *= contribute
		N >>= 1
	}
	return res
}
