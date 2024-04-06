// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/6 20:37
// -----------------------------------------------
package main

// 二叉树

func robIII(root *TreeNode) int {
	var dfs func(*TreeNode) [2]int // robed notRob
	dfs = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}
		l, r := dfs(node.Left), dfs(node.Right)
		robed := node.Val + l[1] + r[1]
		notRob := max(l[0], l[1]) + max(r[0], r[1])
		return [2]int{robed, notRob}
	}
	res := dfs(root)
	return max(res[0], res[1])
}
