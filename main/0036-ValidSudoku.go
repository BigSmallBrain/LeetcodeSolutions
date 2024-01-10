// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/10 21:10
// -----------------------------------------------
package main

// 哈希表

func isValidSudoku(board [][]byte) bool {
	row := [9][9]int{}
	col := [9][9]int{}
	box := [3][3][9]int{}

	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			index := board[i][j] - '1'
			row[i][index]++
			col[j][index]++
			box[i/3][j/3][index]++
			if row[i][index] > 1 || col[j][index] > 1 || box[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
