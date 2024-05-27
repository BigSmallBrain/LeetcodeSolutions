// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/27 12:08
// -----------------------------------------------
package main

// 模拟

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	total := mean * (m + n)
	target := total
	for _, r := range rolls {
		target -= r
	}
	if target < n || target > 6*n { // 不存在
		return []int{}
	}
	res := make([]int, n)
	num, flag := target/n, target%n
	for i := 0; i < n; i++ {
		if i < flag {
			res[i]++
		}
		res[i] += num
	}
	return res
}
