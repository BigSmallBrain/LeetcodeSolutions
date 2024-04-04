// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/4 19:53
// -----------------------------------------------
package main

// 二叉树

func longestUnivaluePath(root *TreeNode) (res int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		l, r := 0, 0
		if node.Left != nil && node.Val == node.Left.Val {
			l += left + 1
		}
		if node.Right != nil && node.Val == node.Right.Val {
			r += right + 1
		}
		res = max(res, l+r)
		return max(l, r)
	}
	dfs(root)
	return
}
