// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/15 13:35
// -----------------------------------------------
package main

import (
	"sort"
)

// 贪心

func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	doTask := map[int]struct{}{}
	for _, task := range tasks {
		start, end, duration := task[0], task[1], task[2]
		// 之前已经完成的任务点
		for i := start; i <= end && duration > 0; i++ {
			if _, ok := doTask[i]; ok {
				duration--
			}
		}
		// 未完成
		for i := end; duration > 0 && i >= start; i-- {
			if _, ok := doTask[i]; !ok {
				doTask[i] = struct{}{}
				duration--
			}
		}
	}
	return len(doTask)
}
