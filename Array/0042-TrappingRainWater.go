// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/23 10:53
// -----------------------------------------------
package main

func trap(height []int) int {
	length := len(height)
	if length == 1 || length == 2 {
		return 0
	}

	// 找到最后一个最大元素索引
	maxIndex, maxHeight := 0, 0
	for i := 0; i < length; i++ {
		if maxHeight <= height[i] {
			maxIndex = i
			maxHeight = height[i]
		}
	}

	// 结果
	sum := 0

	// 左侧
	left := 0
	// 左侧第一个不为0的索引
	for i := 0; i <= maxIndex; i++ {
		if height[i] > 0 {
			left = i
			break
		}
	}

	right := left + 1
	for right <= maxIndex {
		if height[right] >= height[left] {
			sum += trapLeftRain(left, right, height)
			left = right
			right++
		} else {
			right++
		}
	}

	// 右侧
	right = length - 1
	// 右侧第一个不为0的索引
	for i := length - 1; i >= maxIndex; i-- {
		if height[i] > 0 {
			right = i
			break
		}
	}

	left = right - 1
	for left >= maxIndex {
		if height[right] <= height[left] {
			sum += trapRightRain(left, right, height)
			right = left
			left--
		} else {
			left--
		}
	}
	return sum
}

func trapLeftRain(left, right int, height []int) int {
	res := 0
	for i := left + 1; i < right; i++ {
		res += height[left] - height[i]
	}
	return res
}

func trapRightRain(left, right int, height []int) int {
	res := 0
	for i := right - 1; i > left; i-- {
		res += height[right] - height[i]
	}
	return res
}
