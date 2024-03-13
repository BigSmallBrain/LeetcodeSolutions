// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/13 11:09
// -----------------------------------------------
package main

import "strings"

func maximumOddBinaryNumber(s string) string {
	counter := make([]int, 2)
	for _, ch := range s {
		counter[ch-'0']++
	}
	return strings.Repeat("1", counter[1]-1) + strings.Repeat("0", counter[0]) + "1"
}
