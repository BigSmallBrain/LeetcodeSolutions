// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/25 19:20
// -----------------------------------------------
package main

// 回溯

func maxUniqueSplit(s string) (res int) {
	n := len(s)
	unique := map[string]bool{}
	var backtrack func(int, int)
	backtrack = func(start int, end int) {
		if len(unique)+end-start < res {
			return
		}
		if start == end {
			res = max(res, len(unique))
			return
		}
		for i := start + 1; i <= end; i++ {
			curStr := s[start:i]
			if unique[curStr] != true {
				unique[curStr] = true
				backtrack(i, end)
				delete(unique, curStr)
			}
		}
	}
	backtrack(0, n)
	return
}
