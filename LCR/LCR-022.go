// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/9/4 22:28
// -----------------------------------------------
package main

// 快慢指针

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			access := head
			for access != slow {
				access = access.Next
				slow = slow.Next
			}
			return access
		}
	}
	return nil
}
