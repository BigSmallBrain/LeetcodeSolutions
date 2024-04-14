// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/13 20:42
// -----------------------------------------------
package main

// 链表

func reorderList(head *ListNode) {
	// 寻找链表中点
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	temp := slow.Next
	slow.Next = nil
	// 后半部反转
	var tempReverse *ListNode
	for temp != nil {
		next := temp.Next
		temp.Next = tempReverse
		tempReverse = temp
		temp = next
	}
	// 插入
	for curr := head; tempReverse != nil; curr = curr.Next {
		insert := tempReverse
		tempReverse = tempReverse.Next
		insert.Next = curr.Next
		curr.Next = insert
		curr = curr.Next
	}
	return
}
