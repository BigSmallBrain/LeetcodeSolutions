// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/25 19:47
// -----------------------------------------------
package main

import "strings"

// 回溯

func restoreIpAddresses(s string) []string {
	n := len(s)
	path := make([]int, 0)
	res := make([]string, 0)
	var backtrack func(int)
	backtrack = func(index int) {
		if len(path) == 3 {
			nums := make([]string, 4)
			nums[0] = s[:path[0]]
			nums[1] = s[path[0]:path[1]]
			nums[2] = s[path[1]:path[2]]
			nums[3] = s[path[2]:]
			for _, num := range nums {
				if num[0] == '0' && len(num) > 1 {
					return
				}
				if len(num) > 3 {
					return
				}
				if len(num) == 3 && num > "255" {
					return
				}
			}
			res = append(res, strings.Join(nums, "."))
			return
		}
		for i := index; i < n-1; i++ {
			path = append(path, i+1)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
