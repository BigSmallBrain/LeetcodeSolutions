// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/9/11 22:57
// -----------------------------------------------
package main

// 回溯

func getPermutation(n int, k int) (res string) {

	path := make([]byte, 0)
	selected := make([]bool, n)

	var backtrack func(int)
	backtrack = func(flag int) {
		if k == 0 {
			return
		}
		if flag == n {
			k--
			if k == 0 {
				res = string(path)
			}
			return
		}
		for i := 0; i < n; i++ {
			if !selected[i] {
				ch := byte('0' + i + 1)
				selected[i] = true
				path = append(path, ch)
				backtrack(flag + 1)
				selected[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)
	return
}
