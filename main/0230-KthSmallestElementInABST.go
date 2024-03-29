// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/29 10:54
// -----------------------------------------------
package main

// 二叉搜索树

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k == 1 {
			return root.Val
		}
		k--
		root = root.Right
	}
	return 0
}
