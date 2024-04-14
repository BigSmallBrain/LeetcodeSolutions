// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/13 22:05
// -----------------------------------------------
package main

// 链表

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	reverse := func(node *ListNode) *ListNode {
		var head *ListNode
		curr := node
		for curr != nil {
			next := curr.Next
			curr.Next = head
			head = curr
			curr = next
		}
		return head
	}
	l1, l2 = reverse(l1), reverse(l2)
	dummy := &ListNode{0, nil}
	curr := dummy
	flag := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			for l2 != nil {
				l2.Val += flag
				flag = 0
				if l2.Val >= 10 {
					l2.Val %= 10
					flag++
				}
				curr.Next = &ListNode{l2.Val, nil}
				curr, l2 = curr.Next, l2.Next
			}
			break
		}
		if l2 == nil {
			for l1 != nil {
				l1.Val += flag
				flag = 0
				if l1.Val >= 10 {
					l1.Val %= 10
					flag++
				}
				curr.Next = &ListNode{l1.Val, nil}
				curr, l1 = curr.Next, l1.Next
			}
			break
		}
		digit := l1.Val + l2.Val + flag
		flag = 0
		if digit >= 10 {
			digit %= 10
			flag++
		}
		curr.Next = &ListNode{digit, nil}
		l1, l2, curr = l1.Next, l2.Next, curr.Next
	}
	if flag == 1 {
		curr.Next = &ListNode{1, nil}
	}
	// 反转结果
	return reverse(dummy.Next)
}
