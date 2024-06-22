// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/22 14:43
// -----------------------------------------------
package main

// 贪心 字符串

func smallestBeautifulString(s string, k int) string {
	n := len(s)
	res := []byte(s)
	index := -1
	// 修改的第一个字母
	for i := n - 1; i >= 0; i-- {
		for ch := res[i] + 1; ch < 'a'+uint8(k); ch++ {
			if (i > 0 && ch == res[i-1]) || (i > 1 && ch == res[i-2]) {
				continue
			}
			res[i] = ch
			break
		}
		if res[i] != s[i] {
			index = i + 1
			break
		}
	}
	if index < 0 { // 特殊情况
		return ""
	}
	// 修改后续字母
	for i := index; i < n; i++ {
		for ch := uint8('a'); ch < 'a'+uint8(k); ch++ {
			if (i > 0 && ch == res[i-1]) || (i > 1 && ch == res[i-2]) {
				continue
			}
			res[i] = ch
			break
		}
	}
	return string(res)
}
