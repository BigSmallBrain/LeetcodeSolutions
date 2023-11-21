// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/21 21:04
// -----------------------------------------------
package main

import "sort"

// 二分查找

var guess func(int) int

func guessNumber(n int) int {
	return sort.Search(n, func(x int) bool {
		return guess(x) <= 0
	})
}
