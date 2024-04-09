// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/9 10:20
// -----------------------------------------------
package main

// 图

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	// 构建邻接表，所有和 i 节点相连的点都会放入
	adj := make([][]int, n)
	// 记录不同节点邻居的数量
	neighborNum := make([]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		neighborNum[edge[0]]++
		neighborNum[edge[1]]++
	}
	// 初始化队列，包含所有度为 1 的节点（只有 1 个相邻节点的节点，也就是叶子节点）
	var queue []int
	for i := 0; i < n; i++ {
		if neighborNum[i] == 1 {
			queue = append(queue, i)
		}
	}
	// 每层叶子节点加入的结果集
	var result []int
	for len(queue) > 0 {
		// 上一层叶子节点加入的结果集直接清空，我们只需要最后一层的叶子节点即可
		result = []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			// 叶子节点出队列并加进结果集
			leafNode := queue[0]
			queue = queue[1:]
			result = append(result, leafNode)
			neighbors := adj[leafNode]
			for _, neighbor := range neighbors {
				// 每剥离一个叶子节点，这个叶子节点邻居的邻居节点数量减一
				// 当这个叶子节点邻居数量为 1 的时候，外层叶子节点全剥完了，这个邻居节点也成为了叶子节点，加入队列
				neighborNum[neighbor]--
				// adj无需去除相邻节点，因为只判断到 1 就入队列，neighborNum[neighbor]重复减（减到1以下）不影响结果
				if neighborNum[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	return result
}
