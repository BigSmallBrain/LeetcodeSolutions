// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/31 10:20
// -----------------------------------------------
package main

import "strconv"

// 哈希表

func isHappy(n int) bool {
	sumPow := func(num string) int {
		sum := 0
		for i := 0; i < len(num); i++ {
			sum += int(num[i]-'0') * int(num[i]-'0')
		}
		return sum
	}

	condition := func(num int) bool {
		numStr := strconv.Itoa(num)
		if numStr[0] != '1' {
			return false
		}
		for i := 1; i < len(numStr); i++ {
			if numStr[i] != '0' {
				return false
			}
		}
		return true
	}

	flag := map[int]int{}
	cnt := 0
	for cnt <= 1 {
		n = sumPow(strconv.Itoa(n))
		if condition(n) {
			return true
		}
		flag[n]++
		cnt = flag[n]
	}
	return false
}
