// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/26 11:16
// -----------------------------------------------
package main

// 动态规划

func maxProfitIV(k int, prices []int) (res int) {
	n := len(prices)
	buy := make([]int, k)
	sell := make([]int, k)
	for i := 0; i < k; i++ {
		buy[i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], -prices[i])
		sell[0] = max(sell[0], buy[0]+prices[i])
		for j := 1; j < k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	for i := 0; i < k; i++ {
		res = max(res, buy[i], sell[i])
	}
	return
}
