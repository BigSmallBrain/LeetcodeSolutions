// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/29 20:00
// -----------------------------------------------
package main

import (
	"sort"
)

// 双重二分查找

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	maxVal := arr[n-1]

	sum := func(indexVal, xVal int) int {
		ans := 0
		for i := 0; i < indexVal; i++ {
			ans += arr[i]
		}
		ans += (n - indexVal) * xVal
		return ans
	}

	getIndex := func(xVal int) int {
		return sort.Search(n-1, func(i int) bool {
			return arr[i] >= xVal
		})
	}

	res := sort.Search(maxVal, func(x int) bool {
		index := getIndex(x)
		tgt := sum(index, x)
		return tgt > target-1
	})

	if target-sum(getIndex(res-1), res-1) <= sum(getIndex(res), res)-target {
		return res - 1
	}

	return res
}
