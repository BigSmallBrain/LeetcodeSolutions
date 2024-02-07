// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/7 21:06
// -----------------------------------------------
package main

func hardestWorker(n int, logs [][]int) (res int) {
	minVal := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	start, flag := 0, 0
	for _, log := range logs {
		if flag == log[1]-start {
			res = minVal(res, log[0])
		} else if flag < log[1]-start {
			res = log[0]
			flag = log[1] - start
		}
		start = log[1]
	}
	return
}
