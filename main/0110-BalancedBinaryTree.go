// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/30 9:03
// -----------------------------------------------
package main

import "math"

// 平衡二叉树

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var maxDeep func(*TreeNode) int
	maxDeep = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		return max(maxDeep(node.Left), maxDeep(node.Right)) + 1
	}
	return math.Abs(float64(maxDeep(root.Left)-maxDeep(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
