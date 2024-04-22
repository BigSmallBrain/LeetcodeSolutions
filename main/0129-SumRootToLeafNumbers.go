// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/21 12:40
// -----------------------------------------------
package main

// 二叉树

func sumNumbers(root *TreeNode) (res int) {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}
		totalSum := sum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return totalSum
		}
		return dfs(node.Left, totalSum) + dfs(node.Right, totalSum)
	}
	return dfs(root, 0)
}
