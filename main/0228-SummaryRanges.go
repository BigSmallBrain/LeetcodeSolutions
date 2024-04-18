// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/17 10:25
// -----------------------------------------------
package main

import "fmt"

// 数组

func summaryRanges(nums []int) (res []string) {
	for i := 0; i < len(nums); {
		start := i
		for i += 1; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		}
		s := fmt.Sprintf("%d", nums[start])
		if i > start+1 {
			s = s + "->" + fmt.Sprintf("%d", nums[i-1])
		}
		res = append(res, s)
	}
	return
}
