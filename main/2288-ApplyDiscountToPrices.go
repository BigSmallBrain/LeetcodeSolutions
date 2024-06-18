// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/18 11:09
// -----------------------------------------------
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 字符串 正则匹配

func discountPrices(sentence string, discount int) string {
	res := strings.Split(sentence, " ")
	pattern := `^\$[0-9]+$`
	for i, s := range res {
		if matched, _ := regexp.MatchString(pattern, s); matched {
			price, _ := strconv.Atoi(s[1:])
			res[i] = fmt.Sprintf("$%.2f", float64(price)*(1-float64(discount)/100))
		}
	}
	return strings.Join(res, " ")
}
