// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/27 11:00
// -----------------------------------------------
package main

import "slices"

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	value := preorder[0]
	if len(preorder) == 1 {
		return &TreeNode{value, nil, nil}
	}
	index := slices.Index(postorder, preorder[1])
	left := constructFromPrePost(preorder[1:index+2], postorder[:index+1])
	right := constructFromPrePost(preorder[index+2:], postorder[index+1:len(postorder)-1])
	return &TreeNode{value, left, right}
}
