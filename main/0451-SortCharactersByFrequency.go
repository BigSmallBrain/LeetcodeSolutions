// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/3 15:59
// -----------------------------------------------
package main

import (
	"bytes"
	"sort"
)

// 哈希表

func frequencySort(s string) string {
	cnt := map[byte]int{}
	flag := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := cnt[s[i]]; !ok {
			flag = append(flag, s[i])
		}
		cnt[s[i]]++
	}
	sort.Slice(flag, func(i, j int) bool {
		return cnt[flag[i]] > cnt[flag[j]]
	})
	res := make([]byte, 0, len(s))
	for _, ch := range flag {
		res = append(res, bytes.Repeat([]byte{ch}, cnt[ch])...)
	}
	return string(res)
}
