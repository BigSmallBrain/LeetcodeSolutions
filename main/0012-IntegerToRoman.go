// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/2/26 20:44
// -----------------------------------------------
package main

// 哈希表

func intToRoman(num int) (res string) {
	valueSymbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			res += vs.symbol
		}
		if num == 0 {
			break
		}
	}
	return
}
