// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/2 21:24
// -----------------------------------------------
package main

// 二分查找

func shipWithinDays(weights []int, days int) int {
	n := len(weights)

	count := func(x int) int {
		flag := 1
		sum := 0
		for i := 0; i < n; i++ {
			if sum+weights[i] > x {
				flag++
				sum = 0
			}
			sum += weights[i]
		}
		return flag
	}

	l, r := -1, 0
	for i := 0; i < n; i++ {
		if l < weights[i] {
			l = weights[i]
		}
		r += weights[i]
	}

	for l < r {
		mid := (l + r) >> 1
		if count(mid) > days {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
