// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/16 12:52
// -----------------------------------------------
package main

// 单调栈

func removeOuterParentheses(s string) string {
	cnt := 0
	buffer := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if cnt > 0 {
				buffer = append(buffer, s[i])
			}
			cnt++
		}
		if s[i] == ')' {
			if cnt > 1 {
				buffer = append(buffer, s[i])
			}
			cnt--
		}
	}
	return string(buffer)
}
