// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/18 21:49
// -----------------------------------------------
package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = struct{}{}
	}
	return false
}
