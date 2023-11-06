// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/6 15:35
// -----------------------------------------------
package main

import "sort"

// 回溯
func permute(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)

	res := make([][]int, 0)
	state := make([]int, 0)
	selected := make([]bool, n)

	var backtrack func()
	backtrack = func() {
		// 状态长度等于元素数量时，记录解
		if len(state) == n {
			res = append(res, append([]int{}, state...))
			return
		}

		for i := 0; i < n; i++ {
			choice := nums[i]
			if !selected[i] { // 未被选中
				selected[i] = true
				state = append(state, choice)
				backtrack()
				selected[i] = false
				state = state[:len(state)-1]
			}

		}
	}
	backtrack()
	return res
}
