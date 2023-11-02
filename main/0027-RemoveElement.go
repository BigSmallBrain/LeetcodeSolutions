// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/24 16:27
// -----------------------------------------------
package main

// 双指针
func removeElement(nums []int, val int) int {
	flag := len(nums)
	for i := 0; i < flag; {
		if nums[i] == val {
			for j := i; j < flag-1; j++ {
				nums[j] = nums[j+1]
			}
			flag--
		} else {
			i++
		}
	}
	return flag
}
