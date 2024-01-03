// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/3 16:49
// -----------------------------------------------
package main

// 哈希表

func groupAnagrams(strs []string) [][]string {
	flagMap := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, s := range str {
			cnt[s-'a']++
		}
		flagMap[cnt] = append(flagMap[cnt], str)
	}
	res := make([][]string, 0)
	for _, v := range flagMap {
		res = append(res, v)
	}
	return res
}
