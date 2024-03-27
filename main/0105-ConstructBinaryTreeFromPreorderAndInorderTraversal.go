// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/27 10:33
// -----------------------------------------------
package main

import "slices"

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	value := preorder[0]
	index := slices.Index(inorder, value)
	left := buildTree(preorder[1:index+1], inorder[:index])
	right := buildTree(preorder[index+1:], inorder[index+1:])
	return &TreeNode{value, left, right}
}
