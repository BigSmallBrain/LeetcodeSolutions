// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/1 22:36
// -----------------------------------------------
package main

// 哈希表

func isIsomorphic(s string, t string) bool {
	n := len(s)
	srcMap, tgtMap := make([]byte, 256), make([]byte, 256)
	for i := 0; i < n; i++ {
		if srcMap[s[i]] == 0 {
			srcMap[s[i]] = t[i]
		} else if srcMap[s[i]] != t[i] {
			return false
		}
		if tgtMap[t[i]] == 0 {
			tgtMap[t[i]] = s[i]
		} else if tgtMap[t[i]] != s[i] {
			return false
		}
	}
	return true
}
