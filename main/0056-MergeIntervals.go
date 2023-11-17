// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/17 22:13
// -----------------------------------------------
package main

import (
	"math"
	"sort"
)

// 善用res
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		if len(res) == 0 {
			res = append(res, intervals[i])
		} else if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = int(math.Max(float64(res[len(res)-1][1]), float64(intervals[i][1])))
		}
	}
	return res
}
