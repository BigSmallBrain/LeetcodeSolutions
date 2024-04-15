// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/15 19:28
// -----------------------------------------------
package main

import "sort"

// 排序

func hIndex(citations []int) (res int) {
	n := len(citations)
	sort.Ints(citations)
	for i := n - 1; i >= 0; i-- {
		if citations[i] >= n-i {
			res = n - i
		} else {
			break
		}
	}
	return max(0, res)
}
