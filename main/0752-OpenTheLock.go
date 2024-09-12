// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/9/12 17:32
// -----------------------------------------------
package main

// 广度优先搜索

func openLock(deadEnds []string, target string) int {

	start := "0000"

	if start == target {
		return 0
	}

	dead := map[string]bool{}
	for _, end := range deadEnds {
		dead[end] = true
	}

	bfs := func(status string) (ans []string) {
		chs := []byte(status)
		for i, ch := range chs {
			chs[i] = ch - 1
			if chs[i] < '0' {
				chs[i] = '9'
			}
			ans = append(ans, string(chs))
			chs[i] = ch + 1
			if chs[i] > '9' {
				chs[i] = '0'
			}
			ans = append(ans, string(chs))
			chs[i] = ch
		}
		return
	}

	type Pair struct {
		status string
		step   int
	}

	queue := []Pair{
		{start, 0},
	}

	reached := map[string]bool{}
	reached[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if dead[curr.status] {
			continue
		}
		for _, next := range bfs(curr.status) {
			if !reached[next] && !dead[next] {
				if next == target {
					return curr.step + 1
				}
				reached[next] = true
				queue = append(queue, Pair{next, curr.step + 1})
			}
		}
	}
	return -1
}
