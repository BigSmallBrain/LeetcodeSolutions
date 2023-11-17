// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/13 19:16
// -----------------------------------------------
package main

import (
	"sort"
	"strconv"
)

// 自定义排序
func findRelativeRanks(score []int) []string {
	n := len(score)
	tempMap := make(map[int]int)
	for i := 0; i < n; i++ {
		tempMap[score[i]] = i
	}
	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	res := make([]string, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			res[tempMap[score[i]]] = "Gold Medal"
		} else if i == 1 {
			res[tempMap[score[i]]] = "Silver Medal"
		} else if i == 2 {
			res[tempMap[score[i]]] = "Bronze Medal"
		} else {
			res[tempMap[score[i]]] = strconv.Itoa(i + 1)
		}
	}
	return res
}
