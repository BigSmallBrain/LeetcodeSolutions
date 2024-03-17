// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/17 19:31
// -----------------------------------------------
package main

func pathSum(root *TreeNode, targetSum int) (res [][]int) {
	var dfs func(*TreeNode, []int, int)
	dfs = func(r *TreeNode, currList []int, preValue int) {
		if r == nil {
			return
		}
		currValue := r.Val + preValue
		currList = append(currList, r.Val)
		if r.Left == nil && r.Right == nil && currValue == targetSum {
			res = append(res, append([]int{}, currList...))
			return
		}
		dfs(r.Left, currList, currValue)
		dfs(r.Right, currList, currValue)
	}
	dfs(root, []int{}, 0)
	return res
}
