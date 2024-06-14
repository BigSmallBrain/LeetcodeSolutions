// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/14 10:53
// -----------------------------------------------
package main

import (
	"container/heap"
	"sort"
)

// 优先队列

type hp2 struct {
	counter map[int]int
	sort.IntSlice
}

func (h *hp2) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp2) Pop() any {
	last := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return last
}

func (h *hp2) Less(i, j int) bool {
	return h.counter[h.IntSlice[i]] < h.counter[h.IntSlice[j]]
}

func topKFrequent(nums []int, k int) []int {
	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}
	h := &hp2{counter, make(sort.IntSlice, 0)}
	heap.Init(h)
	for num := range counter {
		heap.Push(h, num)
		if len(h.IntSlice) > k {
			heap.Pop(h)
		}
	}
	return h.IntSlice
}
