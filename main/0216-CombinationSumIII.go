// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/26 10:52
// -----------------------------------------------
package main

// 回溯
func combinationSum3(k int, n int) [][]int {
	N := 9
	res := make([][]int, 0)
	path := make([]int, 0)
	sum := 0

	var backtrack func(int)
	backtrack = func(flag int) {
		if len(path) == k && sum == n {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := flag; i < N; i++ {
			choice := i + 1
			path = append(path, choice)
			sum += choice
			backtrack(i + 1)
			sum -= choice
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
