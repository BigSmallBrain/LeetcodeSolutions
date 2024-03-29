// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/29 10:12
// -----------------------------------------------
package main

// 二叉搜索树

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Right != nil {
			temp := root.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			root.Val = temp.Val
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			temp := root.Left
			for temp.Right != nil {
				temp = temp.Right
			}
			root.Val = temp.Val
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}
