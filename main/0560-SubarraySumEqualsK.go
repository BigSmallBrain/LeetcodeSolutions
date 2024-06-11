// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/10 11:12
// -----------------------------------------------
package main

// 前缀和 哈希表

func subarraySum(nums []int, k int) (res int) {
	n, sum := len(nums), 0
	memo := map[int]int{0: 1}
	for i := 0; i < n; i++ {
		sum += nums[i]
		if counter, ok := memo[sum-k]; ok {
			res += counter
		}
		memo[sum]++
	}
	return
}
