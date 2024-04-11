// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/11 10:44
// -----------------------------------------------
package main

// 链表

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	for temp := head; list1 != nil || list2 != nil; temp = temp.Next {
		if list1 == nil {
			temp.Next = list2
			break
		}
		if list2 == nil {
			temp.Next = list1
			break
		}
		if list1.Val < list2.Val {
			temp.Next = &ListNode{list1.Val, nil}
			list1 = list1.Next
		} else {
			temp.Next = &ListNode{list2.Val, nil}
			list2 = list2.Next
		}
	}
	return head.Next
}
