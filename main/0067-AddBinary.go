// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/22 22:16
// -----------------------------------------------
package main

import "strconv"

// 字符串 数学

func addBinary(a string, b string) (res string) {
	m, n := len(a), len(b)
	carry := 0
	length := max(m, n)
	for i := 0; i < length; i++ {
		if i < m {
			carry += int(a[m-i-1] - '0')
		}
		if i < n {
			carry += int(b[n-i-1] - '0')
		}
		res = strconv.Itoa(carry%2) + res
		carry /= 2
	}
	if carry > 0 {
		return "1" + res
	}
	return
}
