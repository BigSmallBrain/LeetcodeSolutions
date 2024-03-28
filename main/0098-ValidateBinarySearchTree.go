// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/28 10:05
// -----------------------------------------------
package main

import "math"

// 二叉搜索树

func isValidBST(root *TreeNode) bool {
	var helper func(*TreeNode, int, int) bool
	helper = func(node *TreeNode, upper int, lower int) bool {
		if node == nil {
			return true
		}
		if node.Val >= upper || node.Val <= lower {
			return false
		}
		return helper(node.Left, node.Val, lower) && helper(node.Right, upper, node.Val)
	}
	return helper(root, math.MaxInt, math.MinInt)
}
