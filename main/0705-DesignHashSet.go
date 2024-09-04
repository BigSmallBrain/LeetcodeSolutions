// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/24 22:19
// -----------------------------------------------
package main

import (
	"container/list"
)

const base = 769

type MyHashSet struct {
	data []list.List
}

func ConstructorMyHashSet() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}

func (s *MyHashSet) Hash(key int) int {
	return key % base
}

func (s *MyHashSet) Add(key int) {
	if !s.Contains(key) {
		h := s.Hash(key)
		s.data[h].PushBack(key)
	}
}

func (s *MyHashSet) Remove(key int) {
	h := s.Hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			s.data[h].Remove(e)
		}
	}
}

func (s *MyHashSet) Contains(key int) bool {
	h := s.Hash(key)
	for e := s.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}
