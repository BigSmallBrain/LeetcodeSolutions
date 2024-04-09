// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/8 20:50
// -----------------------------------------------
package main

func titleToNumber(columnTitle string) (res int) {
	mapping := func(ch byte) int {
		return int(ch-'A') + 1
	}
	flag := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += mapping(columnTitle[i]) * flag
		flag *= 26
	}
	return
}
