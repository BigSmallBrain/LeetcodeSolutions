// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/9 20:18
// -----------------------------------------------
package main

func rotate(matrix [][]int) {
	n := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		l, r := 0, n-1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}
