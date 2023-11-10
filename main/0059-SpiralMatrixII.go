// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/10 14:27
// -----------------------------------------------
package main

func generateMatrix(n int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}

	left, right := 0, n-1
	up, down := 0, n-1
	flag := 1
	for left < right && up < down {
		for i := left; i < right; i++ {
			res[up][i] = flag
			flag++
		}
		for i := up; i < down; i++ {
			res[i][right] = flag
			flag++
		}
		for i := right; i > left; i-- {
			res[down][i] = flag
			flag++
		}
		for i := down; i > up; i-- {
			res[i][left] = flag
			flag++
		}
		left++
		right--
		up++
		down--
	}
	if left == right {
		for i := up; i <= down; i++ {
			res[i][left] = flag
			flag++
		}
		return res
	}
	if up == down {
		for i := left; i <= right; i++ {
			res[up][i] = flag
			flag++
		}
		return res
	}

	return res
}
