// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/12 16:01
// -----------------------------------------------
package main

// 滑动窗口 前缀和

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)

	leftSum := make([]int, k+1)
	for i := 1; i < k+1; i++ {
		leftSum[i] = leftSum[i-1] + cardPoints[i-1]
	}

	rightSum := make([]int, k+1)
	for i, flag := k-1, n-1; i >= 0; i, flag = i-1, flag-1 {
		rightSum[i] = rightSum[i+1] + cardPoints[flag]
	}

	res := 0
	for i := 0; i < k+1; i++ {
		temp := leftSum[i] + rightSum[i]
		if res < temp {
			res = temp
		}
	}
	return res
}
