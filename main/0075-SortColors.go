// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/23 19:20
// -----------------------------------------------
package main

// 双指针
func sortColors(nums []int) {
	l, r := 0, len(nums)-1

	for i := 0; i <= r; i++ {
		for ; i <= r && nums[i] == 2; r-- {
			nums[i], nums[r] = nums[r], nums[i]
		}

		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
}
