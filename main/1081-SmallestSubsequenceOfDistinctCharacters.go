// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/14 14:32
// -----------------------------------------------
package main

import "strings"

// 栈 贪心

func smallestSubsequence(s string) string {
	flag := [26]int{}
	for _, ch := range s {
		flag[ch-'a']++
	}

	stack := make([]byte, 0)
	for _, ch := range s {
		flag[ch-'a']--
		if strings.Contains(string(stack), string(ch)) {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > byte(ch) && flag[stack[len(stack)-1]-'a'] > 0 {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, byte(ch))
	}
	return string(stack)
}
