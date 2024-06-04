// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/4 16:06
// -----------------------------------------------
package main

import "strings"

// 单调栈

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	n := len(num)
	for i := 0; i < n; i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	res := string(stack[:len(stack)-k])
	res = strings.TrimLeft(res, "0")
	if len(res) == 0 {
		return "0"
	}
	return res
}
