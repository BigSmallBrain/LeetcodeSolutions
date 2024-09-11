// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/9/11 22:11
// -----------------------------------------------
package main

import "container/heap"

// 最小堆

type minHeap []*ListNode

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(v any) {
	*m = append(*m, v.(*ListNode))
}

func (m *minHeap) Pop() any {
	last := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	hp := minHeap{}
	for _, l := range lists {
		if l != nil {
			hp = append(hp, l)
		}
	}
	heap.Init(&hp)
	dummy := &ListNode{}
	curr := dummy
	for len(hp) > 0 {
		curr.Next = heap.Pop(&hp).(*ListNode)
		curr = curr.Next
		if curr.Next != nil {
			heap.Push(&hp, curr.Next)
		}
	}
	return dummy.Next
}
