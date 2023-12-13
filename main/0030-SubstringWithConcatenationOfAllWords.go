// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/5 17:04
// -----------------------------------------------
package main

// 滑动窗口 哈希表

func findSubstring(s string, words []string) []int {
	m, n := len(words), len(words[0])
	length := m * n
	if length > len(s) {
		return []int{}
	}

	cmpMap := func(m1, m2 map[string]int) bool {
		if len(m1) != len(m2) {
			return false
		}
		for k, v1 := range m1 {
			v2, ok := m2[k]
			if !ok || v2 != v1 {
				return false
			}
		}
		return true
	}

	wordsCnt := make(map[string]int)
	for i := 0; i < m; i++ {
		wordsCnt[words[i]]++
	}

	res := make([]int, 0)
	for l, r := 0, length; r <= len(s); l, r = l+1, r+1 {
		tempCnt := make(map[string]int)
		for i := l; i < r; i += n {
			tempCnt[s[i:i+n]]++
		}
		if cmpMap(tempCnt, wordsCnt) {
			res = append(res, l)
		}
	}
	return res
}
