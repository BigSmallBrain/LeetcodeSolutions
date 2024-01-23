// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/23 21:33
// -----------------------------------------------
package main

// 动态规划

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j > 0 && j < i {
				row[j] = res[i-1][j] + res[i-1][j-1]
			} else {
				row[j] = 1
			}
		}
		res = append(res, row)
	}
	return res
}
