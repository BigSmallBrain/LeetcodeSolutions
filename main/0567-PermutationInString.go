// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/13 19:42
// -----------------------------------------------
package main

// 滑动窗口 array

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}

	b1 := [26]byte{}
	b2 := [26]byte{}
	for i := 0; i < m; i++ {
		b1[s1[i]-'a']++
		b2[s2[i]-'a']++
	}
	if b1 == b2 {
		return true
	}
	for i := m; i < n; i++ {
		b2[s2[i]-'a']++
		b2[s2[i-m]-'a']--
		if b1 == b2 {
			return true
		}
	}
	return false
}
