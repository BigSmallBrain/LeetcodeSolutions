// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/11 10:29
// -----------------------------------------------
package main

// 链表排序 归并

func sortList(head *ListNode) *ListNode {
	merge := func(head1, head2 *ListNode) *ListNode {
		dummy := &ListNode{}
		temp, temp1, temp2 := dummy, head1, head2
		for temp1 != nil && temp2 != nil {
			if temp1.Val < temp2.Val {
				temp.Next = temp1
				temp1 = temp1.Next
			} else {
				temp.Next = temp2
				temp2 = temp2.Next
			}
			temp = temp.Next
		}
		if temp1 != nil {
			temp.Next = temp1
		} else if temp2 != nil {
			temp.Next = temp2
		}
		return dummy.Next
	}
	var sort func(*ListNode, *ListNode) *ListNode
	sort = func(head *ListNode, tail *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		if head.Next == tail {
			head.Next = nil
			return head
		}
		slow, fast := head, head
		for fast != tail && fast.Next != tail {
			slow = slow.Next
			fast = fast.Next.Next
		}
		mid := slow
		return merge(sort(head, mid), sort(mid, tail))
	}
	return sort(head, nil)
}
