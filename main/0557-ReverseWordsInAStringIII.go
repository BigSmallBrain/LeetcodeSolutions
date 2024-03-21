// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/20 22:01
// -----------------------------------------------
package main

import (
	"strings"
)

func reverseWords(s string) string {
	splits := make([]string, 0)
	for _, str := range strings.Split(s, " ") {
		temp := []byte(str)
		for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
			temp[i], temp[j] = temp[j], temp[i]
		}
		splits = append(splits, string(temp))
	}
	return strings.Join(splits, " ")
}
