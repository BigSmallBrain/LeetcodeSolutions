// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/8 9:46
// -----------------------------------------------
package main

// 贪心

func minimumPossibleSum(n int, target int) (res int) {
	mod := int(1e9 + 7)
	flag := target / 2
	if flag >= n {
		return n * (n + 1) / 2 % mod
	} else {
		return (flag*(flag+1)/2 + (n-flag)*(2*target+n-flag-1)/2) % mod
	}
}
