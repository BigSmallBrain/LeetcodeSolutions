// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/5 17:38
// -----------------------------------------------
package main

import "math"

func findRestaurant(list1 []string, list2 []string) (res []string) {
	m1 := map[string]int{}
	for i, str := range list1 {
		m1[str] = i
	}
	flag := math.MaxInt
	for j, str := range list2 {
		if i, ok := m1[str]; ok {
			if i+j < flag {
				flag = i + j
				res = []string{str}
			} else if i+j == flag {
				res = append(res, str)
			}
		}
	}
	return
}
