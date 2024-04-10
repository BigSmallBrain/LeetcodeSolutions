// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/10 18:37
// -----------------------------------------------
package main

// 链表

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	length := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		length++
	}
	flag := length - k%length
	if flag == length {
		return head
	}
	iter.Next = head
	for flag > 0 {
		iter = iter.Next
		flag--
	}
	res := iter.Next
	iter.Next = nil
	return res
}
