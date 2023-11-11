// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/9 20:46
// -----------------------------------------------
package main

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])

	reverse := func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	a, b := 0, 0
	x, y := 0, 0

	res := make([]int, 0)
	for s := 0; s < m+n-1; s++ {
		temp := make([]int, 0)
		for i, j := a, b; i >= x && j <= y; i, j = i-1, j+1 {
			temp = append(temp, mat[i][j])
		}
		if s%2 == 1 {
			reverse(temp)
		}
		res = append(res, temp...)
		if a < m-1 {
			a++
		} else {
			b++
		}
		if y < n-1 {
			y++
		} else {
			x++
		}
	}

	return res
}
