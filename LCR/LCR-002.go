// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/10/10 12:43
// -----------------------------------------------
package main

func addBinary(a string, b string) string {

	res := make([]byte, 0)

	flag := 0
	m, n := len(a), len(b)
	for i := 0; i < m || i < n; i++ {

		var x, y int
		if m-i-1 >= 0 {
			x = int(a[m-i-1] - '0')
		}
		if n-i-1 >= 0 {
			y = int(b[n-i-1] - '0')
		}

		temp := x + y + flag
		if temp > 1 {
			flag = 1
		} else {
			flag = 0
		}

		res = append(res, byte('0'+temp%2))
	}

	if flag > 0 {
		res = append(res, '1')
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
