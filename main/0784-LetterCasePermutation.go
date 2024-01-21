// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/21 18:45
// -----------------------------------------------
package main

// 回溯

func letterCasePermutation(s string) []string {
	n := len(s)
	res := make([]string, 0)
	path := []byte(s)
	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			res = append(res, string(path))
			return
		}
		backtrack(i + 1)
		if path[i] >= 'A' {
			path[i] ^= 32
			backtrack(i + 1)
		}
	}
	backtrack(0)
	return res
}
