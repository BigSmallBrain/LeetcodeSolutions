// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/12 15:44
// -----------------------------------------------
package main

// 动态规划 背包问题

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		counter := make([]int, 2)
		for j := 0; j < len(str); j++ {
			counter[str[j]-'0']++
		}
		for j := m; j >= 0; j-- {
			for k := n; k >= 0; k-- {
				if j >= counter[0] && k >= counter[1] {
					dp[j][k] = max(dp[j][k], dp[j-counter[0]][k-counter[1]]+1)
				}
			}
		}
	}
	return dp[m][n]
}
