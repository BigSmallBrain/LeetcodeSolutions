// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/18 21:35
// -----------------------------------------------
package main

// 滑动窗口 哈希

func totalFruit(fruits []int) int {
	n := len(fruits)
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fruitMap := map[int]int{}
	res := 0
	l, r := 0, 0
	for r < n {
		fruitMap[fruits[r]]++
		for len(fruitMap) > 2 {
			fruitMap[fruits[l]]--
			if fruitMap[fruits[l]] == 0 {
				delete(fruitMap, fruits[l])
			}
			l++
		}
		res = maxVal(res, r-l+1)
		r++
	}
	return res
}
