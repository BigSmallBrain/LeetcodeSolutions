// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/1 20:08
// -----------------------------------------------
package main

// 动态规划

func maxCoins(nums []int) int {
	values := append([]int{1}, append(nums, 1)...)
	n := len(values)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			maxCoin := -1
			for k := i + 1; k < j; k++ {
				maxCoin = max(maxCoin, values[i]*values[k]*values[j]+dp[i][k]+dp[k][j])
			}
			dp[i][j] = maxCoin
		}
	}
	return dp[0][n-1]
}
