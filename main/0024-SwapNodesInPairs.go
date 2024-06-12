// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/12 13:36
// -----------------------------------------------
package main

// 递归

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := head.Next
	head.Next = swapPairs(h.Next)
	h.Next = head
	return h
}
