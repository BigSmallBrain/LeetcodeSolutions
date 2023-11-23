// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/23 20:08
// -----------------------------------------------
package main

// 二分查找

var isBadVersion func(int) bool

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		temp := (l + r) >> 1
		if isBadVersion(temp) {
			r = temp
		} else {
			l = temp + 1
		}
	}
	return l
}
