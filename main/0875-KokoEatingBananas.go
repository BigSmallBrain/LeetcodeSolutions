// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/1 18:52
// -----------------------------------------------
package main

// 二分查找

func minEatingSpeed(piles []int, h int) int {
	n := len(piles)

	getHours := func(x int) int {
		sum := 0
		for i := 0; i < n; i++ {
			if piles[i]%x == 0 {
				sum += piles[i] / x
			} else {
				sum += piles[i]/x + 1
			}
		}
		return sum
	}

	maxVal := -1
	for i := 0; i < n; i++ {
		if maxVal < piles[i] {
			maxVal = piles[i]
		}
	}
	l, r := 1, maxVal
	for l < r {
		mid := (l + r) >> 1
		if getHours(mid) > h {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
