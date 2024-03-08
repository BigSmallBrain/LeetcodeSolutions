// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/19 19:23
// -----------------------------------------------
package main

// 动态规划

func longestArithSeqLength(nums []int) (res int) {
	minVal, maxVal := 999, -1
	for _, num := range nums {
		minVal = min(minVal, num)
		maxVal = max(maxVal, num)
	}
	diff := maxVal - minVal
	dp := make([][]int, maxVal+1)
	for i := 0; i < maxVal+1; i++ {
		dp[i] = make([]int, 2*diff+1)
	}
	for _, num := range nums {
		for j := 0; j < 2*diff+1; j++ {
			k := num - j + diff
			if k >= minVal && k <= maxVal {
				dp[num][j] = dp[k][j] + 1
			} else {
				dp[num][j] = 1
			}
			res = max(res, dp[num][j])
		}
	}
	return
}
