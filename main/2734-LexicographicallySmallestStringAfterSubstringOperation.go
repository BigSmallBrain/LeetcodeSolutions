// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/27 16:04
// -----------------------------------------------
package main

// 贪心

func smallestString(s string) string {
	chs := []byte(s)
	n := len(s)
	done := false
	for i := 0; i < n; i++ {
		if chs[i] == 'a' {
			if done {
				break
			}
			continue
		}
		chs[i]--
		done = true
	}
	if !done {
		chs[n-1] = 'z'
		return string(chs)
	}
	return string(chs)
}
