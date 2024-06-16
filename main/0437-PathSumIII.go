// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/16 16:37
// -----------------------------------------------
package main

// 双重递归

func pathSumIII(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, tgt int) {
		if node == nil {
			return
		}
		tgt -= node.Val
		if tgt == 0 {
			res++
		}
		dfs(node.Left, tgt)
		dfs(node.Right, tgt)
	}

	// 前序遍历
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node, targetSum)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
