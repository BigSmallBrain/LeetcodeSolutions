// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/25 20:38
// -----------------------------------------------
package main

import (
	"sort"
)

func isPerfectSquare(num int) bool {
	x := sort.Search(num, func(i int) bool {
		return i*i >= num
	})
	if x*x == num {
		return true
	}
	return false
}
