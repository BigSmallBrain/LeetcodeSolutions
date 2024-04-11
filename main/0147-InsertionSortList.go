// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/11 11:15
// -----------------------------------------------
package main

// 链表 插入排序

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{-1, head}
	sorted, curr := head, head.Next
	for curr != nil {
		if sorted.Val <= curr.Val { // 大于已排序最后一个元素 无需操作
			sorted = sorted.Next
		} else {
			prev := dummy
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			} // prev.Next 大于 curr
			sorted.Next = curr.Next // 清除原位置
			curr.Next = prev.Next
			prev.Next = curr // 插入
		}
		curr = sorted.Next
	}
	return dummy.Next
}
