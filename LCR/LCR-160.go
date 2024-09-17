// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/9/17 16:11
// -----------------------------------------------
package main

import "container/heap"

// 优先队列

type minHp []int

func (m *minHp) Len() int {
	return len(*m)
}

func (m *minHp) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHp) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHp) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHp) Pop() any {
	last := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return last
}

type maxHp []int

func (m *maxHp) Len() int {
	return len(*m)
}

func (m *maxHp) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHp) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHp) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHp) Pop() any {
	last := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return last
}

type MedianFinder struct {
	left  maxHp
	right minHp
}

func Constructor() MedianFinder {
	return MedianFinder{make(maxHp, 0), make(minHp, 0)}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == this.right.Len() {
		heap.Push(&this.right, num)
		heap.Push(&this.left, heap.Pop(&this.right))
	} else {
		heap.Push(&this.left, num)
		heap.Push(&this.right, heap.Pop(&this.left))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		return float64(this.left[0]+this.right[0]) / 2
	}
	return float64(this.left[0])
}
