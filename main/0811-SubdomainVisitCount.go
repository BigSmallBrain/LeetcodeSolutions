// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/1/4 13:20
// -----------------------------------------------
package main

import (
	"strconv"
	"strings"
)

// 哈希表

func subdomainVisits(cpdomains []string) []string {
	countMap := map[string]int{}
	for _, cpdomain := range cpdomains {
		splits := strings.Split(cpdomain, " ")
		visit, _ := strconv.Atoi(splits[0])
		domain := splits[1]
		for index := strings.Index(domain, "."); index != -1; {
			countMap[domain] += visit
			domain = domain[index+1:]
			index = strings.Index(domain, ".")
		}
		countMap[domain] += visit
	}
	res := make([]string, 0)
	for domain, visit := range countMap {
		res = append(res, strings.Join([]string{strconv.Itoa(visit), domain}, " "))
	}
	return res
}
