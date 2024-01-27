// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/27 16:17
// -----------------------------------------------
package main

// 埃氏筛法

func countPrimes(n int) (res int) {
	notPrime := make([]bool, n)
	if n > 2 {
		res++
	}
	for i := 3; i < n; i += 2 {
		if !notPrime[i] {
			for j := 3; i*j < n; j += 2 {
				notPrime[i*j] = true
			}
			res++
		}
	}
	return
}
