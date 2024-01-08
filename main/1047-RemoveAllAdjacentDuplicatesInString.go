// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/8 20:25
// -----------------------------------------------
package main

// 栈

func removeDuplicatesIII(s string) string {
	n := len(s)
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
