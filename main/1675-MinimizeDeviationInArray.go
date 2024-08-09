// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/8/9 10:38
// -----------------------------------------------
package main

import (
	"container/heap"
	"math"
	"sort"
)

// 最大堆

type hp4 struct {
	sort.IntSlice
}

func (h *hp4) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp4) Pop() any {
	temp := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return temp
}

func (h *hp4) Len() int {
	return len(h.IntSlice)
}

func (h *hp4) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp4) Swap(i, j int) {
	h.IntSlice[i], h.IntSlice[j] = h.IntSlice[j], h.IntSlice[i]
}

func minimumDeviation(nums []int) int {
	h := hp4{}
	minVal := math.MaxInt
	for i := range nums {
		if nums[i]&1 == 1 {
			nums[i] <<= 1
		}
		minVal = min(minVal, nums[i])
		heap.Push(&h, nums[i])
	}

	res := h.IntSlice[0] - minVal
	for h.IntSlice[0]&1 == 0 {
		currMax := heap.Pop(&h).(int)
		minVal = min(minVal, currMax>>1)
		heap.Push(&h, currMax>>1)
		res = min(res, h.IntSlice[0]-minVal)
	}
	return res
}
