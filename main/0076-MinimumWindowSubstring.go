// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/18 9:37
// -----------------------------------------------
package main

import (
	"math"
)

// 滑动窗口

func minWindow(s string, t string) string {
	getIndex := func(x uint8) uint8 {
		if x < 26 {
			return x
		}
		return x%32 + 26
	}
	flag := [52]int{}
	for i := 0; i < len(t); i++ {
		flag[getIndex(t[i]-'A')]++
	}
	contain := func(src, tgt [52]int) bool {
		for i := 0; i < 52; i++ {
			if tgt[i] > 0 && tgt[i] > src[i] {
				return false
			}
		}
		return true
	}
	counter := [52]int{}
	l, r := 0, 0
	L, R := 0, 0
	n, length := len(s), math.MaxInt
	for r < n {
		counter[getIndex(s[r]-'A')]++
		r++
		for l <= r && contain(counter, flag) {
			if r-l < length {
				length = r - l
				L, R = l, l+length
			}
			counter[getIndex(s[l]-'A')]--
			l++
		}
	}
	return s[L:R]
}
