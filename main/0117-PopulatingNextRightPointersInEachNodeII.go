// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/26 10:40
// -----------------------------------------------
package main

func connectII(root *Node) *Node {
	for start := root; start != nil; {
		var last, nextStart *Node
		handle := func(curr *Node) {
			if curr == nil {
				return
			}
			if nextStart == nil {
				nextStart = curr
			}
			if last != nil {
				last.Next = curr
			}
			last = curr
		}
		for n := start; n != nil; n = n.Next {
			handle(n.Left)
			handle(n.Right)
		}
		start = nextStart
	}
	return root
}
