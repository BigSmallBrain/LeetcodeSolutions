// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/14 15:21
// -----------------------------------------------
package main

// 队列 前缀和

func shortestSubarray(nums []int, k int) int {
	minVal := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	n := len(nums)
	preSumArr := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSumArr[i+1] = preSumArr[i] + nums[i]
	}
	res := n + 1
	queue := make([]int, 0)
	for i, curSum := range preSumArr {
		for len(queue) > 0 && curSum-preSumArr[queue[0]] >= k {
			res = minVal(res, i-queue[0])
			queue = queue[1:]
		}
		for len(queue) > 0 && curSum <= preSumArr[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	if res < n+1 {
		return res
	}
	return -1
}
