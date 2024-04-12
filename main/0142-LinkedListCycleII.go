// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/12 19:58
// -----------------------------------------------
package main

// 链表 快慢指针

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			flag := head
			for flag != slow {
				flag = flag.Next
				slow = slow.Next
			}
			return flag
		}
	}
	return nil
}
