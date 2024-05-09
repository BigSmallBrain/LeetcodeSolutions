// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/7 15:58
// -----------------------------------------------
package main

import "strings"

// 回溯

func solveNQueens(n int) (res [][]string) {
	// 验证queen位置
	isValid := func(row, col int, queens []int) bool {
		for i, queenCol := range queens {
			if col == queenCol || row-i == col-queenCol || row-i == queenCol-col {
				return false
			}
		}
		return true
	}

	var backtrack func(int, []int)
	backtrack = func(row int, queens []int) {
		// 结束
		if row == n {
			board := make([]string, n)
			for i := 0; i < n; i++ {
				board[i] = strings.Repeat(".", queens[i]) + "Q" + strings.Repeat(".", n-queens[i]-1)
			}
			res = append(res, board)
		}

		// 开始回溯
		for col := 0; col < n; col++ {
			if isValid(row, col, queens) {
				queens = append(queens, col)
				backtrack(row+1, queens)
				queens = queens[:len(queens)-1]
			}
		}
	}

	backtrack(0, []int{})
	return
}
