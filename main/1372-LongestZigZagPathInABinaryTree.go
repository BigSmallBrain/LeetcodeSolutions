// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/7 10:42
// -----------------------------------------------
package main

// 二叉树

func longestZigZag(root *TreeNode) (res int) {
	var dfs func(*TreeNode) [2]int // left right
	dfs = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}
		l, r := dfs(node.Left), dfs(node.Right)
		var left, right int
		if node.Left != nil {
			left = l[1] + 1
		}
		if node.Right != nil {
			right = r[0] + 1
		}
		res = max(res, left, right)
		return [2]int{left, right}
	}
	dfs(root)
	return
}
