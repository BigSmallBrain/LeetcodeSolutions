// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/17 17:05
// -----------------------------------------------
package main

import "sort"

// 贪心 双指针

func maxProfitAssignment(difficulty []int, profit []int, workers []int) (res int) {
	n := len(difficulty)
	jobs := make([][2]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [2]int{difficulty[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})
	sort.Ints(workers)
	i, best := 0, 0
	for _, worker := range workers {
		for i < n && worker >= jobs[i][0] {
			best = max(best, jobs[i][0])
			i++
		}
		res += best
	}
	return
}
