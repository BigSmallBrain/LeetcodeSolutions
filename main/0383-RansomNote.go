// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/12/28 21:23
// -----------------------------------------------
package main

// 哈希表

func canConstruct(ransomNote string, magazine string) bool {
	src := [26]int{}
	tgt := [26]int{}
	for _, char := range ransomNote {
		tgt[char-'a']++
	}
	for _, char := range magazine {
		src[char-'a']++
	}
	for i := 0; i < 26; i++ {
		if src[i] < tgt[i] {
			return false
		}
	}
	return true
}
