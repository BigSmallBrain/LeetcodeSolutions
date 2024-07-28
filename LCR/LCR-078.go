// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/7/28 10:04
// -----------------------------------------------
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp := l1
			l1 = l1.Next
			curr.Next = temp
			curr = curr.Next
		} else {
			temp := l2
			l2 = l2.Next
			curr.Next = temp
			curr = curr.Next
		}
	}
	for l1 != nil {
		temp := l1
		l1 = l1.Next
		curr.Next = temp
		curr = curr.Next
	}
	for l2 != nil {
		temp := l2
		l2 = l2.Next
		curr.Next = temp
		curr = curr.Next
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	l := mergeKLists(lists[:n/2])
	r := mergeKLists(lists[n/2:])
	return merge(l, r)
}
