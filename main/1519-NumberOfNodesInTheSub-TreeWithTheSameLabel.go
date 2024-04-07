// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/7 11:12
// -----------------------------------------------
package main

// 二叉树

func countSubTrees(n int, edges [][]int, labels string) []int {
	// 构建图的邻接表
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	// 记录每个节点的标签数量
	result := make([]int, n)
	visited := make([]bool, n)

	// 递归函数，统计子树中各个标签的数量
	var dfs func(node int) [26]int
	dfs = func(node int) [26]int {
		visited[node] = true
		var count [26]int
		count[labels[node]-'a'] = 1

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				subCount := dfs(neighbor)
				for i := range count {
					count[i] += subCount[i]
				}
			}
		}
		result[node] = count[labels[node]-'a']
		return count
	}

	dfs(0) // 从根节点开始递归

	return result
}
