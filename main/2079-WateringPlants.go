// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/8 21:58
// -----------------------------------------------
package main

// 模拟

func wateringPlants(plants []int, capacity int) (res int) {
	curr := capacity
	for i, plant := range plants {
		if curr < plant {
			curr = capacity
			res += 2 * i
		}
		curr -= plant
	}
	res += len(plants)
	return
}
