// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/28 21:47
// -----------------------------------------------
package main

import "sort"

// 二分查找

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	size := m * n
	getIndex := func(index int) (int, int) {
		return index / n, index % n
	}

	tgt := sort.Search(size-1, func(i int) bool {
		x, y := getIndex(i)
		return matrix[x][y] >= target
	})

	x, y := getIndex(tgt)
	if tgt < size && matrix[x][y] == target {
		return true
	}
	return false
}
