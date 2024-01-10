// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/9 22:19
// -----------------------------------------------
package main

// 栈

func calculate(s string) int {
	n := len(s)
	stack := make([]int, 0)
	flag := '+'
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			continue
		}
		temp := 0
		for i < n && s[i] >= '0' && s[i] <= '9' {
			temp = 10*temp + int(s[i]-'0')
			i++
		}
		switch flag {
		case '+':
			stack = append(stack, temp)
		case '-':
			stack = append(stack, -temp)
		case '*':
			stack[len(stack)-1] *= temp
		case '/':
			stack[len(stack)-1] /= temp
		}
		if i < n {
			flag = int32(s[i])
		}
	}
	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}
