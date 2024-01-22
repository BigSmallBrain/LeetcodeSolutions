// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/22 21:41
// -----------------------------------------------
package main

import "sort"

// 回溯

func numTilePossibilities(tiles string) int {
	n := len(tiles)
	tileList := []byte(tiles)
	sort.Slice(tileList, func(i, j int) bool {
		return tileList[i] < tileList[j]
	})
	res := 0
	selected := make([]bool, n)
	var backtrack func(int, int)
	backtrack = func(index int, flag int) {
		if index == flag {
			res++
			return
		}
		for i := range tileList {
			if i > 0 && tileList[i] == tileList[i-1] && !selected[i-1] { // 去重
				continue
			}
			if !selected[i] {
				selected[i] = true
				backtrack(index+1, flag)
				selected[i] = false
			}
		}
	}
	for i := 1; i <= n; i++ {
		backtrack(0, i)
	}
	return res
}
