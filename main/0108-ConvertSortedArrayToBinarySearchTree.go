// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/30 9:19
// -----------------------------------------------
package main

// 平衡二叉树

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	flag := len(nums) / 2
	left := sortedArrayToBST(nums[:flag])
	right := sortedArrayToBST(nums[flag+1:])
	return &TreeNode{nums[flag], left, right}
}
