// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/10 12:57
// -----------------------------------------------
package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0)

	left, right := 0, n-1
	up, down := 0, m-1

	for left < right && up < down {
		for i := left; i < right; i++ {
			res = append(res, matrix[up][i])
		}
		for i := up; i < down; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- {
			res = append(res, matrix[down][i])
		}
		for i := down; i > up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		right--
		up++
		down--
	}

	if left == right {
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][left])
		}
		return res
	}
	if up == down {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		return res
	}

	return res
}
