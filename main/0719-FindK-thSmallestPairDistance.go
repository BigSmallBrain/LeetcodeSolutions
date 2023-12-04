// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/3 21:49
// -----------------------------------------------
package main

import "sort"

// 二分查找 排序

func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	return sort.Search(nums[n-1]-nums[0], func(mid int) bool {
		cnt := 0
		for i := 0; i < n; i++ {
			index := sort.SearchInts(nums[:i], nums[i]-mid)
			cnt += i - index
		}
		return cnt >= k
	})
}
