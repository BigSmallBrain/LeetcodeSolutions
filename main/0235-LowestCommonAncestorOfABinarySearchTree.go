// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/29 11:17
// -----------------------------------------------
package main

// 二叉搜索树

func lowestCommonAncestorII(root, p, q *TreeNode) (res *TreeNode) {
	res = root
	for {
		if p.Val < res.Val && q.Val < res.Val {
			res = res.Left
		} else if p.Val > res.Val && q.Val > res.Val {
			res = res.Right
		} else {
			return
		}
	}
}
