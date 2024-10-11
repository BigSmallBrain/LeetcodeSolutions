// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/10/11 14:49
// -----------------------------------------------
package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := [][]int{}
	if n > 2 {
		for i, x := range nums[:n-2] {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			if x+nums[i+1]+nums[i+2] > 0 {
				break
			}
			if x+nums[n-1]+nums[n-2] < 0 {
				continue
			}
			j, k := i+1, n-1
			for j < k {
				s := x + nums[j] + nums[k]
				if s > 0 {
					k--
				} else if s < 0 {
					j++
				} else {
					ans = append(ans, []int{x, nums[j], nums[k]})
					j = j + 1
					for ; j < k && nums[j] == nums[j-1]; j++ {

					}
					k = k - 1
					for ; j < k && nums[k] == nums[k+1]; k-- {

					}
				}
			}
		}
	}

	return ans
}
