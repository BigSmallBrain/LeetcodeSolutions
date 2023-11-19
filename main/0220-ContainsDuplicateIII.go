// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/16 19:38
// -----------------------------------------------
package main

// 桶排序思想

func getIdx(n, v int) int {
	if n >= 0 {
		return n / v
	}
	return (n+1)/v - 1
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	tempMap := make(map[int]int)
	for i, x := range nums {
		idx := getIdx(x, valueDiff+1)
		if _, ok := tempMap[idx]; ok {
			return true
		}
		if y, ok := tempMap[idx+1]; ok && y-x <= valueDiff {
			return true
		}
		if y, ok := tempMap[idx-1]; ok && x-y <= valueDiff {
			return true
		}

		tempMap[idx] = x
		if i >= indexDiff {
			delete(tempMap, getIdx(nums[i-indexDiff], valueDiff+1))
		}
	}
	return false
}
