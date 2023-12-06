// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/6 15:50
// -----------------------------------------------
package main

import "strconv"

// 双指针

func compress(chars []byte) int {
	pos := 0

	index := 0
	cnt := 1
	for index < len(chars) {
		index++
		if index < len(chars) && chars[index] == chars[index-1] {
			cnt++
		} else {
			chars[pos] = chars[index-1]
			pos++
			if cnt == 1 {
				continue
			}
			for _, digit := range strconv.Itoa(cnt) {
				chars[pos] = byte(digit)
				pos++
			}
			cnt = 1
		}
	}
	return pos
}
