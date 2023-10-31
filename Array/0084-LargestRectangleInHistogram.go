// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/31 9:33
// -----------------------------------------------
package main

// TODO 单调栈
func largestRectangleArea(heights []int) int {

	res := -1

	for i := 0; i < len(heights); i++ {
		left, right := i, i

		for left >= 0 {
			if heights[left] >= heights[i] {
				left--
			} else {
				break
			}
		}

		for right < len(heights) {
			if heights[right] >= heights[i] {
				right++
			} else {
				break
			}
		}

		tempRes := (right - left - 1) * heights[i]
		if res < tempRes {
			res = tempRes
		}
	}
	return res
}
