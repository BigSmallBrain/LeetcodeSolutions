// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/8/11 21:00
// -----------------------------------------------
package main

import (
	"fmt"
	"math"
)

// 动态规划 反向

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	get := func(x, y int) int {
		if x >= m || y >= n {
			return math.MaxInt
		}
		return dp[x][y]
	}
	dp[m-1][n-1] = max(1, 1-dungeon[m-1][n-1])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 { // 已重复计算
				continue
			}

			dp[i][j] = max(1, min(get(i+1, j), get(i, j+1))-dungeon[i][j])

		}
	}
	fmt.Println(dp)
	return dp[0][0]
}
