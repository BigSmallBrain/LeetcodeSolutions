// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/15 10:37
// -----------------------------------------------
package main

import "math"

func maxPathSum(root *TreeNode) int {
	res := -math.MaxInt
	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		left := max(dfs(r.Left), 0)
		right := max(dfs(r.Right), 0)
		curr := left + right + r.Val
		res = max(res, curr)
		return max(left, right) + r.Val
	}
	dfs(root)
	return res
}
