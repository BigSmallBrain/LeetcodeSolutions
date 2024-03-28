// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/28 10:50
// -----------------------------------------------
package main

// 二叉搜索树

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
