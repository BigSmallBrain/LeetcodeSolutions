// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/14 19:13
// -----------------------------------------------
package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)

	res := make([][]int, 0)
	path := make([]int, 0)
	selected := make([]bool, n)
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
			if i > 0 && candidates[i] == candidates[i-1] && !selected[i-1] {
				continue
			}
			choice := candidates[i]
			path = append(path, choice)
			selected[i] = true
			sum += choice
			backtrack(i + 1)
			path = path[:len(path)-1]
			selected[i] = false
			sum -= choice
		}
	}

	backtrack(0)
	return res
}
