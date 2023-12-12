// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/12 15:20
// -----------------------------------------------
package main

// 滑动窗口

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	total := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}

	increase := 0
	for i := 0; i < minutes; i++ {
		increase += grumpy[i] * customers[i]
	}

	res := increase
	for i := minutes; i < n; i++ {
		increase = increase + grumpy[i]*customers[i] - grumpy[i-minutes]*customers[i-minutes]
		if increase > res {
			res = increase
		}
	}
	return res + total
}
