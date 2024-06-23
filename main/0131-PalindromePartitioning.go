// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/23 19:44
// -----------------------------------------------
package main

// 回溯

func partitionII(s string) (res [][]string) {
	n := len(s)
	isPalindrome := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	path := make([]string, 0)
	var backtrack func(int)
	backtrack = func(start int) {
		if start == n {
			// 结束
			res = append(res, append([]string{}, path...))
			return
		}
		for i := start; i < n; i++ {
			if isPalindrome(start, i) {
				path = append(path, s[start:i+1])
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)
	return
}
