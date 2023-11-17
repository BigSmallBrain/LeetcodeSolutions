// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/15 21:25
// -----------------------------------------------
package main

import "sort"

// todo 自定义比较函数

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m, n := len(arr1), len(arr2)
	sort.Ints(arr1)
	state := make([]bool, m)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr1[j] == arr2[i] && !state[j] {
				res = append(res, arr1[j])
				state[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		if !state[i] {
			res = append(res, arr1[i])
		}
	}

	return res
}
