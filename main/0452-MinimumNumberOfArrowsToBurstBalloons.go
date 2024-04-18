// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/18 11:19
// -----------------------------------------------
package main

import "sort"

// 贪心 数组

func findMinArrowShots(points [][]int) int {
	res := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			res++
		} else {
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}
	return res
}
