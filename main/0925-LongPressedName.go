// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/6 16:38
// -----------------------------------------------
package main

// 分离双指针

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0

	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}

	return i == len(name)
}
