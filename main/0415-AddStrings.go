// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/9 21:12
// -----------------------------------------------
package main

import "strconv"

// 双指针

func addStrings(num1 string, num2 string) string {
	res := ""
	m, n := len(num1)-1, len(num2)-1
	flag := 0
	for m >= 0 || n >= 0 {
		digit1 := 0
		if m >= 0 {
			digit1 = int(num1[m] - '0')
			m--
		}

		digit2 := 0
		if n >= 0 {
			digit2 = int(num2[n] - '0')
			n--
		}

		temp := digit1 + digit2 + flag
		res = strconv.Itoa(temp%10) + res
		flag = temp / 10
	}

	if flag == 1 {
		return "1" + res
	}
	return res
}
