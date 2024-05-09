// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/9 16:03
// -----------------------------------------------
package main

// 模拟

func minimumRefill(plants []int, capacityA int, capacityB int) (res int) {
	currA, currB := capacityA, capacityB
	for i, j := 0, len(plants)-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			// 最后一个
			if max(currA, currB) < plants[i] {
				res++
			}
		} else {
			if currA < plants[i] {
				res++
				currA = capacityA
			}
			currA -= plants[i]
			if currB < plants[j] {
				res++
				currB = capacityB
			}
			currB -= plants[j]
		}
	}
	return
}
