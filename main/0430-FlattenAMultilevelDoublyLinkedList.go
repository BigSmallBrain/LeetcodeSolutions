// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/10 19:21
// -----------------------------------------------
package main

type DoublyNode struct {
	Val   int
	Prev  *DoublyNode
	Next  *DoublyNode
	Child *DoublyNode
}

// 双向链表

func flattenII(root *DoublyNode) *DoublyNode {
	for curr := root; curr != nil; curr = curr.Next {
		if curr.Child != nil {
			if curr.Next == nil {
				curr.Child, curr.Next = nil, curr.Child
				curr.Next.Prev = curr
				continue
			}
			next := curr.Next
			child := curr.Child
			for child.Next != nil {
				child = child.Next
			}
			curr.Child, curr.Next = nil, curr.Child
			curr.Next.Prev = curr
			child.Next, next.Prev = next, child
		}
	}
	return root
}
