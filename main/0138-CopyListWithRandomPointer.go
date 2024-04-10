// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/10 20:01
// -----------------------------------------------
package main

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

// 链表

func copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}
	oldIndex := map[*RandomNode]int{}
	index := map[int]*RandomNode{}
	h := &RandomNode{head.Val, nil, nil}
	oldIndex[head], index[0] = 0, h
	flag := 1
	for curr, temp := h, head; temp.Next != nil; curr, temp = curr.Next, temp.Next {
		curr.Next = &RandomNode{temp.Next.Val, nil, nil}
		oldIndex[temp.Next], index[flag] = flag, curr.Next
		flag++
	}
	for curr := h; head != nil; curr, head = curr.Next, head.Next {
		if head.Random != nil {
			curr.Random = index[oldIndex[head.Random]]
		}
	}
	return h
}
