// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/22 10:17
// -----------------------------------------------
package main

import "math"

// 二叉树

func minDiffInBST(root *TreeNode) int {
	res, prev := math.MaxInt, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != -1 {
			res = min(res, node.Val-prev)
		}
		prev = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return res
}
