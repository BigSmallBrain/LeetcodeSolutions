// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/12 14:34
// -----------------------------------------------
package main

import "strconv"

// 栈

func evalRPN(tokens []string) int {
	evaluate := func(a, b int, flag string) int {
		switch flag {
		default:
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		}
	}

	stack := make([]int, 0)
	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = evaluate(stack[len(stack)-1], pop, token)
		}
	}
	return stack[0]
}
