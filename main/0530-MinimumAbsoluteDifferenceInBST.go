// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/21 13:22
// -----------------------------------------------
package main

import "math"

// 二叉树

func getMinimumDifference(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return math.MaxInt
	}
	minVal := func(node *TreeNode) (ans int) {
		for temp := node; temp != nil; temp = temp.Left {
			ans = temp.Val
		}
		return
	}
	maxVal := func(node *TreeNode) (ans int) {
		for temp := node; temp != nil; temp = temp.Right {
			ans = temp.Val
		}
		return
	}
	res := math.MaxInt
	if root.Left != nil {
		res = min(res, root.Val-maxVal(root.Left), getMinimumDifference(root.Left))
	}
	if root.Right != nil {
		res = min(res, minVal(root.Right)-root.Val, getMinimumDifference(root.Right))
	}
	return res
}
