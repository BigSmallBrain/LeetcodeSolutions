// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/9 17:32
// -----------------------------------------------
package main

// 回溯

func totalNQueens(n int) (res int) {
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
			res++
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
