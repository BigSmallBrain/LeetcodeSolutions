// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/7 22:04
// -----------------------------------------------
package main

import (
	"strings"
)

// 双指针
func reverseVowels(s string) string {
	res := []byte(s)
	n := len(res)
	l, r := 0, n-1

	for l < r {
		for l < n && !strings.Contains("aeiouAEIOU", string(res[l])) {
			l++
		}
		for r > 0 && !strings.Contains("aeiouAEIOU", string(res[r])) {
			r--
		}
		if l < r {
			res[l], res[r] = res[r], res[l]
			l++
			r--
		}
	}

	return string(res)
}
