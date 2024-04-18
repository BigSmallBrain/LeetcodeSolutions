// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/18 8:45
// -----------------------------------------------
package main

// 数组

func insert(intervals [][]int, newInterval []int) (res [][]int) {
	index := 0
	isOverlap := func(x, y []int) bool {
		return x[0] <= y[1] && x[1] >= y[0]
	}
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			res = append(res, interval)
			index++
		} else {
			break
		}
	}
	res = append(res, newInterval)
	for i := index; i < len(intervals); i++ {
		if isOverlap(res[len(res)-1], intervals[i]) {
			res[len(res)-1][0] = min(intervals[i][0], res[len(res)-1][0])
			res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
