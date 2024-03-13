// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/13 10:54
// -----------------------------------------------
package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	queue := []*TreeNode{root}
	order := make([][]int, 0)
	for len(queue) > 0 {
		level := make([]int, 0)
		for size := len(queue); size > 0; size-- {
			curr := queue[0]
			queue = queue[1:]
			level = append(level, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		if len(level) > 0 {
			order = append(order, level)
		}
	}
	return order
}
