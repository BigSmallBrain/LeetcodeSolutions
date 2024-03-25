// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/25 10:35
// -----------------------------------------------
package main

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	flag := true
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == nil {
			flag = false
		} else {
			if !flag {
				return false
			}
			queue = append(queue, curr.Left, curr.Right)
		}
	}
	return true
}
