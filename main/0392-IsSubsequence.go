// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/17 10:05
// -----------------------------------------------
package main

// 双指针

func isSubsequence(s string, t string) bool {
	index := 0
	for i := 0; i < len(t) && index < len(s); i++ {
		if s[index] == t[i] {
			index++
		}
	}
	if index == len(s) {
		return true
	}
	return false
}
