// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/15 15:01
// -----------------------------------------------
package main

import "strconv"

// 栈

func calPoints(operations []string) int {
	stack := make([]int, 0)
	for _, operation := range operations {
		score, err := strconv.Atoi(operation)
		if err == nil {
			stack = append(stack, score)
			continue
		}
		switch operation {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			stack = stack[:len(stack)-1]
		}
	}
	res := 0
	for _, score := range stack {
		res += score
	}
	return res
}
