// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/8 19:21
// -----------------------------------------------
package main

// 双指针

func backspaceCompare(s string, t string) bool {
	skipS, skipT := 0, 0
	m, n := len(s)-1, len(t)-1

	for m >= 0 || n >= 0 {
		for m >= 0 {
			if s[m] == '#' {
				skipS++
				m--
			} else if skipS > 0 {
				skipS--
				m--
			} else {
				break
			}
		}

		for n >= 0 {
			if t[n] == '#' {
				skipT++
				n--
			} else if skipT > 0 {
				skipT--
				n--
			} else {
				break
			}
		}

		if m >= 0 && n >= 0 {
			if s[m] != t[n] {
				return false
			}
		} else if m >= 0 || n >= 0 {
			return false
		}
		m--
		n--
	}
	return true
}
