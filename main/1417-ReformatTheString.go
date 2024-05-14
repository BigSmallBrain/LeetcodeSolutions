// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/6 21:19
// -----------------------------------------------
package main

import "math"

// 字符串

func reformat(s string) string {
	nums := make([]byte, 0)
	chars := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			nums = append(nums, s[i])
		} else if s[i] >= 'a' && s[i] <= 'z' {
			chars = append(chars, s[i])
		}
	}
	m, n := len(nums), len(chars)
	if math.Abs(float64(m-n)) > 1 {
		return ""
	}
	res := make([]byte, 0)
	if m > n {
		res = append(res, nums[0])
		for i := 0; i < n; i++ {
			res = append(res, chars[i])
			res = append(res, nums[i+1])
		}
	} else if m == n {
		for i := 0; i < n; i++ {
			res = append(res, chars[i])
			res = append(res, nums[i])
		}
	} else {
		res = append(res, chars[0])
		for i := 0; i < m; i++ {
			res = append(res, nums[i])
			res = append(res, chars[i+1])
		}
	}
	return string(res)
}
