// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/15 19:48
// -----------------------------------------------
package main

import (
	"math/rand"
)

// 哈希表

type RandomizedSet struct {
	indices map[int]int
	nums    []int
}

func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{map[int]int{}, []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; ok {
		return false
	}
	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.indices[val]; ok {
		last := this.nums[len(this.nums)-1]
		this.nums[index] = last
		this.indices[last] = index
		this.nums = this.nums[:len(this.nums)-1]
		delete(this.indices, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
