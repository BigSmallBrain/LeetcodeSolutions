// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/6 15:34
// -----------------------------------------------
package main

// 贪心

func minimumSteps(s string) (res int64) {
	n := len(s)
	var counter int64
	for i := 0; i < n; i++ {
		counter += int64(s[i] - '0')
		if s[i] == '0' {
			res += counter
		}
	}
	return
}
