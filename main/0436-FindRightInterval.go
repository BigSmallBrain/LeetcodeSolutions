// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/1 12:18
// -----------------------------------------------
package main

import "sort"

// 二分查找

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	res := make([]int, n)
	indexDict := map[int]int{} // start index
	for i, interval := range intervals {
		indexDict[interval[0]] = i
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		index := sort.Search(n, func(x int) bool {
			return intervals[x][0] >= interval[1]
		})
		i := indexDict[interval[0]]
		if index < n {
			res[i] = indexDict[intervals[index][0]]
		} else {
			res[i] = -1
		}
	}
	return res
}
