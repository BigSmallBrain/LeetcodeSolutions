// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/18 22:18
// -----------------------------------------------
package main

// 栈

type RecentCounter struct {
	requests []int
}

func RecentCounterConstructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)
	start := t - 3000
	for len(this.requests) > 0 && this.requests[0] < start {
		this.requests = this.requests[1:]
	}
	return len(this.requests)
}
