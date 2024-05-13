// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/23 21:02
// -----------------------------------------------
package main

// 背包问题

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum >> 1
	dp := make([]int, target+1)
	for _, stone := range stones {
		for i := target; i >= stone; i-- {
			dp[i] = max(dp[i], dp[i-stone]+stone)
		}
	}
	return sum - 2*dp[target]
}
