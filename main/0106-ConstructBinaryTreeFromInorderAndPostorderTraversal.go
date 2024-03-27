// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/27 10:49
// -----------------------------------------------
package main

import "slices"

func buildTreeII(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	value := postorder[len(postorder)-1]
	index := slices.Index(inorder, value)
	left := buildTreeII(inorder[:index], postorder[:index])
	right := buildTreeII(inorder[index+1:], postorder[index:len(postorder)-1])
	return &TreeNode{value, left, right}
}
