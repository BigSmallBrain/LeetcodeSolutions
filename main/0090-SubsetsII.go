// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/26 14:30
// -----------------------------------------------
package main

import (
	"sort"
)

// 动态规划
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	res := make([][]int, 0)
	temp := make([]int, 0)

	var dfs func(int)
	dfs = func(flag int) {
		res = append(res, append([]int{}, temp...))
		for i := flag; i < n; i++ {
			if i > flag && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0)

	return res
}
