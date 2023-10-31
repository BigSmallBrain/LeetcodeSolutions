// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/31 16:44
// -----------------------------------------------
package main

import "math"

// TODO 单调栈
func sumSubarrayMins(arr []int) int {
	n := int(math.Pow(10, 9)) + 7
	res := 0
	for i := 0; i < len(arr); i++ {
		minVal := arr[i]
		for j := i; j < len(arr); j++ {
			if minVal > arr[j] {
				minVal = arr[j]
			}
			res += minVal
			res %= n
		}
	}
	return res
}
