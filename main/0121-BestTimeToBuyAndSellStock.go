// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/14 18:17
// -----------------------------------------------
package main

import "math"

func maxProfit0121(prices []int) (res int) {
	minPrice := math.MaxInt
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > res {
			res = price - minPrice
		}
	}
	return
}
