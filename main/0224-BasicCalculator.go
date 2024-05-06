// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/6 12:29
// -----------------------------------------------
package main

// 栈

func calculateII(s string) (ans int) {
	stack := []int{1}
	flag := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			flag = stack[len(stack)-1]
			i++
		case '-':
			flag = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, flag)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			num := 0
			for i < n && s[i] >= '0' && s[i] <= '9' {
				num = 10*num + int(s[i]-'0')
				i++
			}
			ans += flag * num
		}
	}
	return
}
