// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/16 22:04
// -----------------------------------------------
package main

// 链表

type listNode struct {
	Val  int
	Next *listNode
}

func deleteDuplicates(head *listNode) *listNode {
	if head == nil {
		return head
	}

	res := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return res
}
