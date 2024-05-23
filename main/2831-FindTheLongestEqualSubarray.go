// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/23 10:18
// -----------------------------------------------
package main

// 哈希表

func longestEqualSubarray(nums []int, k int) (res int) {
	count := map[int]int{}

	left := 0
	for right := 0; right < len(nums); right++ {
		count[nums[right]]++
		for right-left+1-count[nums[left]] > k {
			count[nums[left]]--
			left++
		}
		res = max(res, count[nums[right]])
	}
	return
}
