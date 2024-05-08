// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/1 21:03
// -----------------------------------------------
package main

import (
	"container/heap"
	"sort"
)

var a []int

type hp1 struct {
	sort.IntSlice
}

func (h *hp1) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp1) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp1) Pop() any {
	last := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return last
}

// 优先队列

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	h := &hp1{make([]int, k)}
	for i := 0; i < k; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)

	n := len(nums)
	res := make([]int, 0)
	res = append(res, nums[h.IntSlice[0]])
	for i := k; i < n; i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}
		res = append(res, nums[h.IntSlice[0]])
	}
	return res
}
