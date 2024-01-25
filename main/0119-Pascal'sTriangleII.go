// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/23 21:44
// -----------------------------------------------
package main

// 动态规划

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i < rowIndex+1; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}
