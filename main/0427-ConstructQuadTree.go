// Package main
// -----------------------------------------------
// @author  : Kingslayer
// @school  : NJUST
// @e-mail  : kingslayer5437@gmail.com
// @time    : 2024/4/24 20:45
// -----------------------------------------------
package main

type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}

// 四叉树

func construct(grid [][]int) *QuadNode {
	n := len(grid)
	if n == 1 {
		return &QuadNode{Val: grid[0][0] == 1, IsLeaf: true}
	}
	isSame := func(g *[][]int) bool {
		flag := (*g)[0][0]
		for i := 0; i < len(*g); i++ {
			for j := 0; j < len(*g); j++ {
				if flag != (*g)[i][j] {
					return false
				}
			}
		}
		return true
	}
	if isSame(&grid) {
		return &QuadNode{Val: grid[0][0] == 1, IsLeaf: true}
	}
	topLeft := make([][]int, n/2)
	topRight := make([][]int, n/2)
	bottomLeft := make([][]int, n/2)
	bottomRight := make([][]int, n/2)
	for i, temp := range grid {
		if i < n/2 {
			topLeft[i] = temp[:n/2]
			topRight[i] = temp[n/2:]
		} else {
			bottomLeft[i%(n/2)] = temp[:n/2]
			bottomRight[i%(n/2)] = temp[n/2:]
		}
	}
	return &QuadNode{
		IsLeaf:      false,
		TopLeft:     construct(topLeft),
		TopRight:    construct(topRight),
		BottomLeft:  construct(bottomLeft),
		BottomRight: construct(bottomRight),
	}
}
