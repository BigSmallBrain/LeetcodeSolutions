// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/14 18:46
// -----------------------------------------------
package main

import "sort"

// 回溯
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)
	sum := 0

	var backtrack func(int)
	backtrack = func(index int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		} else if sum > target {
			return
		}

		for i := index; i < n; i++ {
			if i > 0 && candidates[i] == candidates[i-1] {
				continue
			}
			choice := candidates[i]
			path = append(path, choice)
			sum += choice
			backtrack(i)
			path = path[:len(path)-1]
			sum -= choice
		}
	}
	backtrack(0)
	return res
}
