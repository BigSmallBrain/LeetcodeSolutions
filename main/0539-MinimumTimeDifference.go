// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/8/8 22:04
// -----------------------------------------------
package main

import (
	"sort"
)

// 模拟

func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	sort.Strings(timePoints)

	countMinutes := func(t string) int {
		return (int(t[0]-'0')*10+int(t[1]-'0'))*60 + int(t[3]-'0')*10 + int(t[4]-'0')
	}

	prev := countMinutes(timePoints[0])
	res := 24*60 + prev - countMinutes(timePoints[n-1])
	for i := 1; i < n; i++ {
		curr := countMinutes(timePoints[i])
		res = min(res, curr-prev)
		prev = curr
	}
	return res
}
