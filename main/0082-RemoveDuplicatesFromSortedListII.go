// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/26 18:43
// -----------------------------------------------
package main

// 链表

func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			x := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == x {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}
