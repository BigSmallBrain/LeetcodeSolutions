// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/5/11 18:21
// -----------------------------------------------
package main

import (
	"fmt"
	"strings"
)

// 贪心

func fullJustify(words []string, maxWidth int) (res []string) {
	n := len(words)
	i := 0
	for i < n {
		width := 0
		start := i
		for width <= maxWidth && i < n {
			if width == 0 {
				width += len(words[i])
				i++
			} else if width+len(words[i])+1 <= maxWidth {
				width += len(words[i]) + 1
				i++
			} else {
				break
			}
		}

		gaps := i - start - 1
		if i == n || gaps == 0 { // 单个单词或最后一行
			line := strings.Join(words[start:i], " ")
			res = append(res, line+strings.Repeat(" ", maxWidth-len(line)))
			continue
		}
		// 划分空格
		spaces := maxWidth - width + gaps
		flag, index := spaces/gaps, spaces%gaps
		line := words[start]
		for j := 1; j <= gaps; j++ {
			if j <= index {
				line += strings.Repeat(" ", flag+1) + words[start+j]
			} else {
				line += strings.Repeat(" ", flag) + words[start+j]
			}
		}
		res = append(res, line)
	}
	fmt.Println(res)
	return
}
