// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/19 12:24
// -----------------------------------------------
package main

// 回溯

func combine(n int, k int) (res [][]int) {
	path := make([]int, 0)
	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return
}
