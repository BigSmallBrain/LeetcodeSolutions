// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/17 12:32
// -----------------------------------------------
package main

// 双指针 字符串

func findLUSlength(strs []string) int {
	isSubstr := func(a, b string) bool {
		i := 0
		for j := 0; j < len(b) && i < len(a); j++ {
			if a[i] == b[j] {
				i++
			}
		}
		return i == len(a)
	}
	res := -1
	for i, s := range strs {
		cur := -1
		for j, str := range strs {
			if i != j {
				if isSubstr(s, str) {
					cur = -1
					break
				} else {
					cur = max(cur, len(s))
				}
			}
		}
		res = max(res, cur)
	}
	return res
}
