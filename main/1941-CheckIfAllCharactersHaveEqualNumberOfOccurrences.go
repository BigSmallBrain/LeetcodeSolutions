// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/27 22:45
// -----------------------------------------------
package main

import "sort"

func areOccurrencesEqual(s string) bool {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	sort.Ints(cnt)
	flag := cnt[len(cnt)-1]
	for i := 0; i < 26; i++ {
		if cnt[i] != 0 && flag != cnt[i] {
			return false
		}
	}
	return true
}
