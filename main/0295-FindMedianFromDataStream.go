// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/19 13:02
// -----------------------------------------------
package main

import (
	"container/heap"
	"sort"
)

// 最小堆 最大堆

type MedianFinder struct {
	left, right hp3
}

type hp3 struct {
	sort.IntSlice
}

func (h *hp3) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp3) Pop() any {
	last := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return last
}

func MedianFinderConstructor() MedianFinder {
	return MedianFinder{}
}

func (f *MedianFinder) AddNum(num int) {
	l, r := &f.left, &f.right
	if r.Len() == 0 || num > r.IntSlice[0] {
		heap.Push(r, num)
		if r.Len() > l.Len()+1 {
			heap.Push(l, -heap.Pop(r).(int))
		}
	} else {
		heap.Push(l, -num)
		if r.Len() < l.Len() {
			heap.Push(r, -heap.Pop(l).(int))
		}
	}
}

func (f *MedianFinder) FindMedian() float64 {
	l, r := f.left, f.right
	if l.Len() < r.Len() {
		return float64(r.IntSlice[0])
	}
	return float64(-l.IntSlice[0]+r.IntSlice[0]) / 2
}
