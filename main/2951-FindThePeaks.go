// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/25 21:21
// -----------------------------------------------
package main

// 枚举

func findPeaks(mountain []int) (res []int) {
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			res = append(res, i)
		}
	}
	return
}
