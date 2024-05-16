// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/16 10:52
// -----------------------------------------------
package main

// 贪心

func numberOfWeeks(milestones []int) int64 {
	var total int64
	longest := -1 // 耗时最长工作
	for _, milestone := range milestones {
		longest = max(longest, milestone)
		total += int64(milestone)
	}
	rest := total - int64(longest)
	if int64(longest) <= rest+1 { // 完成所有工作
		return total
	} else { // 无法完成所有
		return 2*rest + 1
	}
}
