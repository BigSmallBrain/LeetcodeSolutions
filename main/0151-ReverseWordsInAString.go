// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/16 21:04
// -----------------------------------------------
package main

import (
	"slices"
	"strings"
)

// 字符串

func reverseWordsII(s string) string {
	splits := strings.Split(s, " ")
	res := make([]string, 0)
	for _, split := range splits {
		if split == "" || split == " " {
			continue
		}
		res = append(res, split)
	}
	slices.Reverse(res)
	return strings.Join(res, " ")
}
