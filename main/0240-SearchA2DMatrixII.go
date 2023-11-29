// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/28 22:06
// -----------------------------------------------
package main

// 类似二叉搜索树

func searchMatrixII(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if target > matrix[i][j] {
			i++
		} else if target < matrix[i][j] {
			j--
		} else {
			return true
		}
	}
	return false
}
