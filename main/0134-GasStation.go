// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/16 22:10
// -----------------------------------------------
package main

// 贪心

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	currGas := 0
	sumGas := 0
	start := 0
	for i := 0; i < n; i++ {
		currGas += gas[i] - cost[i]
		sumGas += gas[i] - cost[i]
		if currGas < 0 {
			start = i + 1
			currGas = 0
		}
	}
	if sumGas < 0 {
		return -1
	}
	return start
}
