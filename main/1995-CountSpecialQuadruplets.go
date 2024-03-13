// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/13 10:25
// -----------------------------------------------
package main

// 枚举

func countQuadruplets(nums []int) (res int) {
	n := len(nums)
	counter := map[int]int{}
	for b := n - 3; b >= 1; b-- {
		for _, d := range nums[b+2:] {
			counter[d-nums[b+1]]++
		}
		for _, a := range nums[:b] {
			if y := a + nums[b]; counter[y] > 0 {
				res += counter[y]
			}
		}
	}
	return
}
