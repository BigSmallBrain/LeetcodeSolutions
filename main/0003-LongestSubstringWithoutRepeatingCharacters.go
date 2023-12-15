// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/15 21:44
// -----------------------------------------------
package main

// 滑动窗口

func lengthOfLongestSubstring(s string) int {
	m := [256]bool{}
	res := 0
	l := 0
	for r, v := range s {
		for m[v] {
			m[s[l]] = false
			l++
		}
		m[v] = true
		if res < r-l+1 {
			res = r - l + 1
		}
	}
	return res
}
