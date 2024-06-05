// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/5 20:58
// -----------------------------------------------
package main

import (
	"strconv"
	"strings"
)

// 字符串

func compareVersion(version1 string, version2 string) int {
	v1s := strings.Split(version1, ".")
	v2s := strings.Split(version2, ".")
	n := max(len(v1s), len(v2s))
	for i := 0; i < n; i++ {
		v1, v2 := 0, 0
		if i < len(v1s) {
			v1, _ = strconv.Atoi(v1s[i])
		}
		if i < len(v2s) {
			v2, _ = strconv.Atoi(v2s[i])
		}
		if v1 > v2 {
			return 1
		}
		if v2 > v1 {
			return -1
		}
	}
	return 0
}
