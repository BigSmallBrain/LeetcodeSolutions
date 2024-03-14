// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/14 10:49
// -----------------------------------------------
package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	queue := []*TreeNode{root}
	order := make([][]int, 0)
	flag := 0
	for len(queue) > 0 {
		level := make([]int, 0)
		for i := len(queue); i > 0; i-- {
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
			if flag%2 == 1 {
				for l, r := 0, len(level)-1; l < r; l, r = l+1, r-1 {
					level[l], level[r] = level[r], level[l]
				}
			}
			order = append(order, level)
		}
		flag++
	}
	return order
}
