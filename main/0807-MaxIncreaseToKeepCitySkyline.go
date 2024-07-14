// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/14 18:44
// -----------------------------------------------
package main

import "slices"

func maxIncreaseKeepingSkyline(grid [][]int) (res int) {
	m, n := len(grid), len(grid[0])
	row := make([]int, m)
	col := make([]int, n)
	for i := 0; i < m; i++ {
		row[i] = slices.Max(grid[i])
		for j := 0; j < n; j++ {
			col[j] = max(col[j], grid[i][j])
		}
	}
	// 计数
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += min(row[i], col[j]) - grid[i][j]
		}
	}
	return
}
