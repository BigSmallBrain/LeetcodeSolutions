// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/23 22:09
// -----------------------------------------------
package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	return len(s) - strings.LastIndex(s, " ") - 1
}
