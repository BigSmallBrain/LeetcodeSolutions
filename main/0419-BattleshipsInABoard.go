// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/11 13:09
// -----------------------------------------------
package main

// 遍历

func countBattleships(board [][]byte) (res int) {
	for i, row := range board {
		for j, ch := range row {
			if ch == 'X' && !(i > 0 && board[i-1][j] == 'X' || j > 0 && board[i][j-1] == 'X') {
				res++
			}
		}
	}
	return
}
