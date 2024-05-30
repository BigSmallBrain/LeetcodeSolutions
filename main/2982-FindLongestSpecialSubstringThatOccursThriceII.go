// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/30 10:45
// -----------------------------------------------
package main

// 双指针

func maximumLengthII(s string) (res int) {
	n := len(s)
	counter := make([][3]int, 26) // 第一 第二 第三
	for i, j := 0, 0; i < n; i = j {
		for j < n && s[j] == s[i] { // 寻找最长
			j++
		}
		length := j - i
		index := s[i] - 'a'
		if length > counter[index][0] {
			counter[index][0], counter[index][1], counter[index][2] = length, counter[index][0], counter[index][1]
		} else if length > counter[index][1] {
			counter[index][1], counter[index][2] = length, counter[index][1]
		} else if length > counter[index][2] {
			counter[index][2] = length
		}
	}
	for _, c := range counter {
		res = max(res, max(c[0]-2, min(c[0]-1, c[1]), c[2]))
	}
	if res > 0 {
		return res
	}
	return -1
}
