// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/22 17:42
// -----------------------------------------------
package main

import "sort"

// 哈希表

func findWinners(matches [][]int) [][]int {
	freq := map[int]int{}
	for _, match := range matches {
		winner, loser := match[0], match[1]
		if freq[winner] == 0 {
			freq[winner] = 0
		}
		freq[loser]++
	}
	res := make([][]int, 2)
	for player, loss := range freq {
		if loss < 2 {
			res[loss] = append(res[loss], player)
		}
	}
	sort.Ints(res[0])
	sort.Ints(res[1])
	return res
}
