// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/10/19 9:58
// -----------------------------------------------
package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	res := make([][]int, 0)

	for a := 0; a < n-3; a++ {
		x := nums[a]

		if a > 0 && x == nums[a-1] { // 跳过重复
			continue
		}

		if x+nums[a+1]+nums[a+2]+nums[a+3] > target {
			break
		}

		if x+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for b := a + 1; b < n-2; b++ {
			y := nums[b]

			if b > a+1 && y == nums[b-1] {
				continue
			}

			if x+y+nums[b+1]+nums[b+2] > target {
				break
			}

			if x+y+nums[n-2]+nums[n-1] < target {
				continue
			}
			c, d := b+1, n-1
			for d > c {
				sum := x + y + nums[c] + nums[d]
				if sum > target {
					d--
				}

				if sum < target {
					c++
				}

				if sum == target {
					res = append(res, []int{x, y, nums[c], nums[d]})
					for c++; c < d && nums[c] == nums[c-1]; c++ {
					}
					for d--; d > c && nums[d] == nums[d+1]; d-- {
					}
				}
			}
		}
	}
	return res
}
