// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/25 13:27
// -----------------------------------------------
package main

// 链表

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	curr := this.head
	for i := 0; i <= index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	this.size++
	curr := this.head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.Next = &ListNode{val, curr.Next}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	this.size--
	curr := this.head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
}
