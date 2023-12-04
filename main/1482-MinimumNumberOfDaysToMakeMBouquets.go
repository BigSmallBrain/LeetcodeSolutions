// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/3 20:37
// -----------------------------------------------
package main

import (
	"math"
	"sort"
)

// 二分查找

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if n < m*k {
		return -1
	}

	minDay, maxDay := math.MaxInt, -1
	for i := 0; i < n; i++ {
		if minDay > bloomDay[i] {
			minDay = bloomDay[i]
		}
		if maxDay < bloomDay[i] {
			maxDay = bloomDay[i]
		}
	}

	return minDay + sort.Search(maxDay-minDay, func(days int) bool {
		days += minDay
		bouquets, bloomed := 0, 0
		for i := 0; i < n; i++ {
			if bloomDay[i] > days {
				bloomed = 0
			} else {
				bloomed++
				if bloomed == k {
					bouquets++
					bloomed = 0
				}
			}
		}
		return bouquets >= m
	})
}
