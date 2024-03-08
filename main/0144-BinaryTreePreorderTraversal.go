// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/7 12:22
// -----------------------------------------------
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {
	var preorder func(*TreeNode)
	preorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		res = append(res, r.Val)
		preorder(r.Left)
		preorder(r.Right)
	}
	preorder(root)
	return
}
