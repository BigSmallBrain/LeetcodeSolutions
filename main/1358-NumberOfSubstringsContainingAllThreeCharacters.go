// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/22 14:11
// -----------------------------------------------
package main

func numberOfSubstrings(s string) int {
	n := len(s)
	cnt := make([]int, 3)
	l, r := 0, 0
	res := 0
	for l < n-2 {
		for r < n && (cnt[0] < 1 || cnt[1] < 1 || cnt[2] < 1) {
			cnt[s[r]-'a']++
			r++
		}
		if cnt[0] >= 1 && cnt[1] >= 1 && cnt[2] >= 1 {
			res += n - r + 1
		}
		cnt[s[l]-'a']--
		l++
	}
	return res
}
