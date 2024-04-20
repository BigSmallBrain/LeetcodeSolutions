// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/20 18:53
// -----------------------------------------------
package main

// 双向链表

type LRUCache struct {
	cache      map[int]*DLinkedNode
	capacity   int
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func LRUCacheConstructor(capacity int) LRUCache {
	l := LRUCache{
		map[int]*DLinkedNode{},
		capacity,
		&DLinkedNode{0, 0, nil, nil},
		&DLinkedNode{0, 0, nil, nil},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		node = &DLinkedNode{key, value, nil, nil}
		this.addToHead(node)
		this.cache[key] = node
		if len(this.cache) > this.capacity {
			removedKey := this.removeTail()
			delete(this.cache, removedKey)
		}
		return
	}
	node.value = value
	this.moveToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() int {
	key := this.tail.prev.key
	this.removeNode(this.tail.prev)
	return key
}
