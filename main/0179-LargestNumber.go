// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/20 18:46
// -----------------------------------------------
package main

import (
	"sort"
	"strconv"
	"strings"
)

// 字符串比较原理
func largestNumber(nums []int) string {
	n := len(nums)
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		if strs[i][0] == strs[j][0] {
			return strs[i]+strs[j] > strs[j]+strs[i]
		}
		return strs[i] > strs[j]
	})

	res := strings.Join(strs, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}
