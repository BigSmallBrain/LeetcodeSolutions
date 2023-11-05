// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/5 12:35
// -----------------------------------------------
package main

// KMP算法

func generateNext(str string) []int {
	n := len(str)
	next := make([]int, n)

	left := 0
	for right := 1; right < n; right++ {
		if left > 0 && str[left] != str[right] {
			left = next[left-1]
		}
		if str[left] == str[right] {
			left++
		}
		next[right] = left
	}
	return next
}

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)

	next := generateNext(needle)

	j := 0
	for i := 0; i < m; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - j + 1
		}
	}
	return -1
}
