// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/18 20:08
// -----------------------------------------------
package main

// 栈

type MyQueue struct {
	queue []int
}

func MyQueueConstructor() MyQueue {
	return MyQueue{make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyQueue) Pop() int {
	flag := this.queue[0]
	this.queue = this.queue[1:]
	return flag
}

func (this *MyQueue) Peek() int {
	return this.queue[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}
