// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/9/4 21:59
// -----------------------------------------------
package main

import "container/list"

type entry struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func ConstructorMyHashMap() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (s *MyHashMap) Hash(key int) int {
	return key % base
}

func (s *MyHashMap) Put(key int, value int) {
	h := s.Hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if kv := e.Value.(entry); kv.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	s.data[h].PushBack(entry{key, value})
}

func (s *MyHashMap) Get(key int) int {
	h := s.Hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			return e.Value.(entry).value
		}
	}
	return -1
}

func (s *MyHashMap) Remove(key int) {
	h := s.Hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			s.data[h].Remove(e)
		}
	}
}
