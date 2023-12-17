// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/17 13:19
// -----------------------------------------------
package main

// 链表迭代

func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = res
		res = curr
		curr = next
	}
	return res
}
