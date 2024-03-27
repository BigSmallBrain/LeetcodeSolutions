// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/26 12:13
// -----------------------------------------------
package main

func flatten(root *TreeNode) {
	order := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		order = append(order, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	flag := root
	for len(order) > 0 {
		curr := order[0]
		order = order[1:]
		flag.Val = curr
		flag.Left = nil
		if len(order) == 0 {
			break
		}
		flag.Right = &TreeNode{}
		flag = flag.Right
	}
}
