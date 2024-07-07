// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/7 11:17
// -----------------------------------------------
package main

// 枚举

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	n := 8
	isGood := func(dx, dy int) bool {
		x, y := rMove+dx, cMove+dy
		step := 1
		for x >= 0 && x < n && y >= 0 && y < n {
			if step == 1 {
				// 第一个相反
				if board[x][y] == '.' || board[x][y] == color {
					return false
				}
			} else {
				if board[x][y] == '.' {
					return false
				}
				if board[x][y] == color {
					return true
				}
			}
			step++
			x += dx
			y += dy
		}
		return false
	}
	dx := []int{1, 1, 0, -1, -1, -1, 0, 1} // 行改变量
	dy := []int{0, 1, 1, 1, 0, -1, -1, -1} // 列改变量
	for i := 0; i < n; i++ {
		if isGood(dx[i], dy[i]) {
			return true
		}
	}
	return false
}
