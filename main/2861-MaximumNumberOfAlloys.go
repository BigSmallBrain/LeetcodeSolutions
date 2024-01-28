// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/27 17:22
// -----------------------------------------------
package main

import (
	"slices"
	"sort"
)

// 二分查找

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) (res int) {
	mx := slices.Min(stock) + budget
	for _, comp := range composition {
		res += sort.Search(mx-res, func(num int) bool {
			num += res + 1
			money := 0
			for i, s := range stock {
				if s < num*comp[i] {
					money += (num*comp[i] - s) * cost[i]
					if money > budget {
						return true
					}
				}
			}
			return false
		})
	}
	return
}
