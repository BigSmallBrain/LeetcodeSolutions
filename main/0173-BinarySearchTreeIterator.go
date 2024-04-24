// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/24 21:27
// -----------------------------------------------
package main

// 二叉树

type BSTIterator struct {
	stack []int
}

func BSTIteratorConstructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{}
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		preorder(node.Left)
		iter.stack = append(iter.stack, node.Val)
		preorder(node.Right)
	}
	preorder(root)
	return iter
}

func (this *BSTIterator) Next() int {
	next := this.stack[0]
	this.stack = this.stack[1:]
	return next
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
