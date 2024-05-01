// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/1 20:10
// -----------------------------------------------
package main

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	last := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return last
}

func (h *hp) Replace(x int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = x
	heap.Fix(h, 0)
	return top
}

// 最小堆

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	res := 0
	if candidates*2+k > n {
		sort.Ints(costs)
		for i := 0; i < k; i++ {
			res += costs[i]
		}
		return int64(res)
	}
	x, y := candidates-1, n-candidates
	hL := &hp{costs[:x+1]}
	hR := &hp{costs[y:]}
	heap.Init(hL)
	heap.Init(hR)
	for i := 0; i < k; i++ {
		l, r := hL.IntSlice[0], hR.IntSlice[0]
		if l <= r {
			res += l
			x++
			hL.Replace(costs[x])
		} else {
			res += r
			y--
			hR.Replace(costs[y])
		}
	}
	return int64(res)
}
