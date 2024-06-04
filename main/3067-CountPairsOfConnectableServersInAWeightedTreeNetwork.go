// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/6/4 11:10
// -----------------------------------------------
package main

// 深度优先搜索

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	graph := make([][][]int, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		graph[u] = append(graph[u], []int{v, w})
		graph[v] = append(graph[v], []int{u, w})
	}

	var dfs func(int, int, int) int
	dfs = func(p int, root int, curr int) int {
		ans := 0
		if curr == 0 {
			ans++
		}
		for _, e := range graph[p] {
			v, cost := e[0], e[1]
			if v != root {
				ans += dfs(v, p, (curr+cost)%signalSpeed)
			}
		}
		return ans
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		pre := 0
		for _, e := range graph[i] {
			v, cost := e[0], e[1]
			counter := dfs(v, i, cost%signalSpeed)
			res[i] += pre * counter
			pre += counter
		}
	}
	return res
}
