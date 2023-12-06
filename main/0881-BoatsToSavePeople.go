// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/6 15:11
// -----------------------------------------------
package main

import "sort"

// 双指针 贪心

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	n := len(people)

	res := 0
	l, r := 0, n-1
	for l <= r {
		if people[l]+people[r] > limit {
			res++
			r--
		} else {
			res++
			l++
			r--
		}
	}
	return res
}
