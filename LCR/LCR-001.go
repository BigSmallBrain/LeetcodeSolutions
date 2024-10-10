// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/10/10 13:13
// -----------------------------------------------
package main

func divide(dividend int, divisor int) int {
	positive := true
	if dividend < 0 {
		dividend = -dividend
		positive = !positive
	}
	if divisor < 0 {
		divisor = -divisor
		positive = !positive
	}

	if divisor > dividend {
		return 0
	}

	res, mulDivisor := 1, divisor
	for dividend > mulDivisor+mulDivisor {
		res += res
		mulDivisor += mulDivisor
	}

	res += divide(dividend-mulDivisor, divisor)
	if !positive {
		res = -res
	}

	if res > 1<<31-1 || -res > 1<<31 {
		return 1<<31 - 1
	}
	return res
}
