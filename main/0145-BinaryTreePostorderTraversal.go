// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/10 21:36
// -----------------------------------------------
package main

func postorderTraversal(root *TreeNode) (res []int) {
	var postorder func(node *TreeNode)
	postorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		postorder(r.Left)
		postorder(r.Right)
		res = append(res, r.Val)
	}
	postorder(root)
	return
}
