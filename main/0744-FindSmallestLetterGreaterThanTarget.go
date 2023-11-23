// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2023/11/23 20:19
// -----------------------------------------------
package main

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {
	index := sort.Search(len(letters), func(i int) bool {
		return letters[i] >= target+1
	})
	if index == len(letters) {
		return letters[0]
	}
	return letters[index]
}
