// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/20 20:15
// -----------------------------------------------
package main

import "math/rand"

// 洗牌思路

type Solution struct {
	Unshuffled []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.Unshuffled
}

func (this *Solution) Shuffle() []int {
	n := len(this.Unshuffled)
	nums := make([]int, n)
	copy(nums, this.Unshuffled)

	for i := 0; i < n; i++ {
		index := rand.Intn(n - i)
		nums[index], nums[n-i-1] = nums[n-i-1], nums[index]
	}
	return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
