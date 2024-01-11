// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/11 21:09
// -----------------------------------------------
package main

// 栈

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	flag := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[flag] {
			stack = stack[:len(stack)-1]
			flag++
		}
	}
	return len(stack) == 0
}
