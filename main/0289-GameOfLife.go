// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/9 22:20
// -----------------------------------------------
package main

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	getStatus := func(a, b int) (status int) {
		status = 0
		if b-1 >= 0 {
			status += board[a][b-1]
			if a-1 >= 0 {
				status += board[a-1][b-1]
			}
		}
		if b+1 < n {
			status += board[a][b+1]
			if a+1 < m {
				status += board[a+1][b+1]
			}
		}
		if a-1 >= 0 {
			status += board[a-1][b]
			if b+1 < n {
				status += board[a-1][b+1]
			}
		}
		if a+1 < m {
			status += board[a+1][b]
			if b-1 >= 0 {
				status += board[a+1][b-1]
			}
		}
		return
	}

	setZero := func(a, b int) {
		board[a][b] = 0
	}

	setOne := func(a, b int) {
		board[a][b] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			s := getStatus(i, j)
			if board[i][j] == 0 && s == 3 {
				defer setOne(i, j)
			} else if board[i][j] == 1 {
				if s < 2 || s > 3 {
					defer setZero(i, j)
				}
			}
		}
	}
	return
}
