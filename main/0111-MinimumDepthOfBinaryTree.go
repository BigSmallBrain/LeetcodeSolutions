// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/14 20:56
// -----------------------------------------------
package main

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	res := math.MaxInt
	if root.Left != nil {
		res = min(res, minDepth(root.Left))
	}
	if root.Right != nil {
		res = min(res, minDepth(root.Right))
	}
	return res + 1
}
