// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/27 19:16
// -----------------------------------------------
package main

// 图 深度优先搜素

func numIslands(grid [][]byte) (res int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x-1, y) // up
		dfs(x+1, y) // down
		dfs(x, y-1) // left
		dfs(x, y+1) // right
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				res++
			}
		}
	}
	return
}
