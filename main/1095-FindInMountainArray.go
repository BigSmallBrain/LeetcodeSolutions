// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/29 21:05
// -----------------------------------------------
package main

import "sort"

type MountainArray struct {
	arr []int
}

func (this *MountainArray) get(index int) int {
	return this.arr[index]
}
func (this *MountainArray) length() int {
	return len(this.arr)
}

// 二次二分查找

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	index := sort.Search(n-1, func(i int) bool {
		return mountainArr.get(i) > mountainArr.get(i+1)
	})
	l := sort.Search(index-1, func(i int) bool {
		return mountainArr.get(i) >= target
	})
	if l < index && mountainArr.get(l) == target {
		return l
	}
	r := sort.Search(n-index-1, func(i int) bool {
		return mountainArr.get(n-i-1) >= target
	})
	if r <= n-index && mountainArr.get(n-r-1) == target {
		return n - r - 1
	}
	return -1
}
