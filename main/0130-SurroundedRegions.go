// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/27 19:30
// -----------------------------------------------
package main

// 图 深度优先搜索

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(x int, y int) {
		if x < 0 || y < 0 || x >= m || y >= n || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'o'
		dfs(x-1, y) // up
		dfs(x+1, y) // down
		dfs(x, y-1) // left
		dfs(x, y+1) // right
	}
	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}
	for i := 1; i < m-1; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'o' {
				board[i][j] = 'O'
			}
		}
	}
}
