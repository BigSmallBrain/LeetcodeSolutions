// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/14 10:50
// -----------------------------------------------
package main

// 贪心

func minimumRounds(tasks []int) (res int) {
	counter := map[int]int{}
	for _, task := range tasks {
		counter[task]++
	}
	for _, nums := range counter {
		if nums < 2 {
			return -1
		}
		round := nums / 3
		if nums%3 != 0 {
			round++
		}
		res += round
	}
	return
}
