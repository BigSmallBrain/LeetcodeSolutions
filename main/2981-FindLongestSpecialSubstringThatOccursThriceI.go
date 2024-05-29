// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/29 12:23
// -----------------------------------------------
package main

import "sort"

// 哈希表 双指针 排序

func maximumLength(s string) int {
	n := len(s)
	counter := map[byte][]int{}
	for i := 0; i < n; i++ {
		counter[s[i]] = append(counter[s[i]], i)
	}
	res := -1
	for _, indices := range counter {
		if len(indices) >= 3 {
			res = max(res, 1)
			// 判断最长连续索引
			length := len(indices)
			temp := make([]int, 0)
			left, right := 0, 1
			for right < length {
				for right < length && indices[left]+right == indices[right]+left {
					temp = append(temp, right-left+1)
					right++
				}
				left = right
				right++
			}
			if len(temp) > 2 {
				sort.Slice(temp, func(i, j int) bool {
					return temp[i] > temp[j]
				})
				res = max(res, temp[2])
			}
		}
	}
	return res
}
