// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/20 12:29
// -----------------------------------------------
package main

func countBeautifulPairs(nums []int) (res int) {
	for i, x := range nums {
		for x >= 10 {
			x /= 10
		}
		for _, y := range nums[i+1:] {
			y %= 10
			if gcd(x, y) == 1 {
				res++
			}
		}
	}
	return
}

// 最大法公约数
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
