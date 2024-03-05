// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/3/4 10:41
// -----------------------------------------------
package main

// 动态规划 单调栈

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	solution := func() int64 {
		left, right := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			right[i] = n
		}
		stack := []int{-1}
		for i, num := range nums {
			for len(stack) > 1 && nums[stack[len(stack)-1]] <= num {
				right[stack[len(stack)-1]] = i
				stack = stack[:len(stack)-1]
			}
			left[i] = stack[len(stack)-1]
			stack = append(stack, i)
		}
		ans := 0
		for i, num := range nums {
			ans += (i - left[i]) * (right[i] - i) * num
		}
		return int64(ans)
	}
	res := solution()
	for i := 0; i < n; i++ {
		nums[i] = -nums[i]
	}
	return res + solution()
}
