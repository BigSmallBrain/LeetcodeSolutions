// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/26 21:15
// -----------------------------------------------
package main

// 回溯

func subsets(nums []int) (res [][]int) {
	n := len(nums)
	path := make([]int, 0)
	var backtrack func(int)
	backtrack = func(index int) {
		res = append(res, append([]int{}, path...))
		if index == n {
			return
		}
		for i := index; i < n; i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}
