// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/13 16:13
// -----------------------------------------------
package main

// 滑动窗口

func maxVowels(s string, k int) int {
	n := len(s)
	vowelMap := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	cnt := 0
	for i := 0; i < k; i++ {
		if _, ok := vowelMap[s[i]]; ok {
			cnt++
		}
	}

	res := cnt
	for index := k; index < n; index++ {
		if _, ok := vowelMap[s[index]]; ok {
			cnt++
		}
		if _, ok := vowelMap[s[index-k]]; ok {
			cnt--
		}
		if cnt > res {
			res = cnt
		}
	}

	return res
}
