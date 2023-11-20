// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/19 17:46
// -----------------------------------------------
package main

func containsNearbyDuplicate(nums []int, k int) bool {
	tempMap := map[int]int{}
	for i, num := range nums {
		if index, ok := tempMap[num]; ok && k >= i-index {
			return true
		}
		tempMap[num] = i
	}
	return false
}
