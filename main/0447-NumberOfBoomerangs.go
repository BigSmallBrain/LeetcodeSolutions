// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/4 14:22
// -----------------------------------------------
package main

// 哈希表

func numberOfBoomerangs(points [][]int) int {
	flag := map[int]int{}
	res := 0
	for _, p1 := range points {
		clear(flag)
		for _, p2 := range points {
			distance2 := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
			res += flag[distance2] * 2
			flag[distance2]++
		}
	}
	return res
}
