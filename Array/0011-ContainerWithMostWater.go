// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/20 9:33
// -----------------------------------------------
package main

func maxArea(height []int) int {

	left, right := 0, len(height)-1
	maxContainer := 0

	for right > left {
		if height[right] <= height[left] {
			if maxContainer < (right-left)*height[right] {
				maxContainer = (right - left) * height[right]
			}
			right--
		} else {
			if maxContainer < (right-left)*height[left] {
				maxContainer = (right - left) * height[left]
			}
			left++
		}
	}
	return maxContainer
}
