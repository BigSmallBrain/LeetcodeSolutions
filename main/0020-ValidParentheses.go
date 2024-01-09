// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/9 21:55
// -----------------------------------------------
package main

// 栈

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	flagMap := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if len(stack) > 0 && flagMap[stack[len(stack)-1]] == s[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}
