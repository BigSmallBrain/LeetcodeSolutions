// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/28 21:15
// -----------------------------------------------
package main

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	splits := strings.Split(s, " ")
	if len(splits) != len(pattern) {
		return false
	}
	ch2word := map[byte]string{}
	word2ch := map[string]byte{}
	for i, split := range splits {
		flag := pattern[i]
		if (ch2word[flag] != "" && ch2word[flag] != split) || (word2ch[split] >= 'a' && word2ch[split] != flag) {
			return false
		}
		ch2word[flag] = split
		word2ch[split] = flag
	}
	return true
}
