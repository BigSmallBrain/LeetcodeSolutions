// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/15 21:12
// -----------------------------------------------
package main

// 滑动窗口

func characterReplacement(s string, k int) int {
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	bytesMap := [26]int{}
	l := 0
	res := 1
	maxCnt := 0
	for r, v := range s {
		bytesMap[v-'A']++
		maxCnt = maxVal(maxCnt, bytesMap[v-'A'])
		for l <= r && (r-l+1)-maxCnt > k {
			bytesMap[s[l]-'A']--
			l++
		}
		res = maxVal(res, r-l+1)
		r++
	}
	return res
}
