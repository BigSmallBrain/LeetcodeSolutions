// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/13 17:18
// -----------------------------------------------
package main

// 双指针

func partitionLabels(s string) (res []int) {
	n := len(s)
	counter := make([]int, 26)
	for i := 0; i < n; i++ {
		counter[s[i]-'a'] = i
	}
	l, r := 0, 0
	for i, ch := range s {
		r = max(r, counter[ch-'a'])
		if i == r {
			res = append(res, r-l+1)
			l = r + 1
		}
	}
	return
}
