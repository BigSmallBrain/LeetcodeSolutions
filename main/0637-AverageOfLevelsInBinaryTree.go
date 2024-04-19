// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/19 11:52
// -----------------------------------------------
package main

// 树

func averageOfLevels(root *TreeNode) (res []float64) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		levelSum := 0
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			levelSum += curr.Val
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		res = append(res, float64(levelSum)/float64(size))
	}
	return
}
