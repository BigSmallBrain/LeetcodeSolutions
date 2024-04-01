// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/1 10:45
// -----------------------------------------------
package main

// 位运算

func grayCode(n int) (res []int) {
	if n == 1 {
		return []int{0, 1}
	}
	res = grayCode(n - 1)
	size := len(res)
	carry := 1 << (n - 1)
	for i := size - 1; i >= 0; i-- {
		res = append(res, res[i]|carry)
	}
	return
}
