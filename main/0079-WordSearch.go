// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/25 9:35
// -----------------------------------------------
package main

// 回溯

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var isExist func(int, int, int) bool
	isExist = func(i, j, flag int) bool {
		if flag == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if used[i][j] || board[i][j] != word[flag] {
			return false
		}
		used[i][j] = true
		find := isExist(i+1, j, flag+1) || isExist(i-1, j, flag+1) || isExist(i, j+1, flag+1) || isExist(i, j-1, flag+1)

		if find {
			return true
		} else {
			used[i][j] = false
			return false
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && isExist(i, j, 0) {
				return true
			}
		}
	}
	return false
}
