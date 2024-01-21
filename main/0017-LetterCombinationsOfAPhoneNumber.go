// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/21 18:20
// -----------------------------------------------
package main

// 回溯

func letterCombinations(digits string) []string {
	mapping := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	n := len(digits)
	res := make([]string, 0)
	var backtrack func(int, []byte)
	backtrack = func(index int, letter []byte) {
		if index == n {
			if len(letter) > 0 {
				res = append(res, string(letter))
			}
			return
		}
		doMapping := mapping[digits[index]-'2']
		for i := 0; i < len(doMapping); i++ {
			letter = append(letter, doMapping[i])
			backtrack(index+1, letter)
			letter = letter[:len(letter)-1]
		}
	}
	backtrack(0, make([]byte, 0))
	return res
}
