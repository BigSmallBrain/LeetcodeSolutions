// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/21 22:16
// -----------------------------------------------
package main

// 数组

func getWinner(arr []int, k int) int {
	maxVal := arr[0]
	win := 0
	for i := 1; i < len(arr) && win < k; i++ {
		if arr[i] > maxVal { // 新的最大值
			maxVal = arr[i]
			win = 0
		}
		win++
	}
	return maxVal
}
