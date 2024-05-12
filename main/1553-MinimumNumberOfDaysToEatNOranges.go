// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/12 20:59
// -----------------------------------------------
package main

// 动态规划 不等式

func minDaysII(n int) int {
	counter := map[int]int{}
	counter[0], counter[1] = 0, 1
	var helper func(int) int
	helper = func(x int) int {
		if days, ok := counter[x]; ok {
			return days
		}
		counter[x] = 1 + min(x%2+helper(x/2), x%3+helper(x/3))
		return counter[x]
	}
	return helper(n)
}
