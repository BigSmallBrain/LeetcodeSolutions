// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/12 13:52
// -----------------------------------------------
package main

// 递归

func reverseKGroup(head *ListNode, k int) *ListNode {
	last := head
	flag := k
	for last != nil && flag > 0 {
		last = last.Next
		flag--
	} // last为递归头节点
	if flag > 0 {
		return head
	}
	dummy := &ListNode{}
	curr := head
	for curr != last { // 翻转
		temp := curr.Next
		curr.Next = dummy.Next
		dummy.Next = curr
		curr = temp
	}
	head.Next = reverseKGroup(last, k)
	return dummy.Next
}
