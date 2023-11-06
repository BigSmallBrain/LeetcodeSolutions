// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/24 14:46
// -----------------------------------------------
package main

// 回溯
func subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)

	var dfs func(int2 int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int{}, path...))
			return
		}

		// 不选nums[i]
		dfs(i + 1)

		// 选nums[i]
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}

	dfs(0)
	return res
}
