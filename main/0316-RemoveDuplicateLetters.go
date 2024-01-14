// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/13 22:23
// -----------------------------------------------
package main

import "strings"

// 栈 贪心

func removeDuplicateLetters(s string) string {
	stack := ""
	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		if strings.Contains(stack, ch) {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > s[i] && strings.Contains(s[i+1:], string(stack[len(stack)-1])) {
			stack = stack[:len(stack)-1]
		}
		stack += ch
	}
	return stack
}
