// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/11 21:43
// -----------------------------------------------
package main

import (
	"strings"
)

// 栈 字符串

func simplifyPath(path string) string {
	stack := make([]string, 0)
	for _, str := range strings.Split(path, "/") {
		if str == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if str != "" && str != "." {
			stack = append(stack, str)
		}
	}
	return "/" + strings.Join(stack, "/")
}
