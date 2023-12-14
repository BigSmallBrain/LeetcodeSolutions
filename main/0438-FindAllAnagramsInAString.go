// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/14 12:01
// -----------------------------------------------
package main

// 滑动窗口

func findAnagrams(s string, p string) []int {
	n, window := len(s), len(p)
	if window > n {
		return []int{}
	}
	src, tgt := [26]byte{}, [26]byte{}
	for i := 0; i < window; i++ {
		src[s[i]-'a']++
		tgt[p[i]-'a']++
	}
	res := make([]int, 0)
	if src == tgt {
		res = append(res, 0)
	}
	for index := window; index < len(s); index++ {
		src[s[index-window]-'a']--
		src[s[index]-'a']++
		if src == tgt {
			res = append(res, index-window+1)
		}
	}
	return res
}
