// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/22 22:05
// -----------------------------------------------
package main

// 动态规划 回溯

func generateParenthesis(n int) []string {
	res := make([][]string, n+1)
	res[0] = append(res[0], "")
	res[1] = append(res[1], "()")
	for flag := 2; flag <= n; flag++ {
		for i := 0; i <= flag-1; i++ {
			for _, s1 := range res[i] {
				for _, s2 := range res[flag-i-1] {
					res[flag] = append(res[flag], "("+s1+")"+s2)
				}
			}
		}
	}
	return res[n]
}
