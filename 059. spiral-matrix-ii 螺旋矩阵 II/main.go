package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(generateMatrix(3)) // [[1,2,3],[8,9,4],[7,6,5]]
	helper.PrintJSON(generateMatrix(1)) // [[1]]
}

// func generateMatrix(n int) [][]int {
// 	total := n * n
// 	res := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		res[i] = make([]int, n)
// 	}
// 	round := 0
// 	i := 1
// 	for round <= n / 2 {
// 		// 上
// 		for row, col := round, round; col < n - round; col++ {
// 			res[row][col] = i
// 			i++
// 		}
// 		if i > total {
// 			break
// 		}
// 		// 右
// 		for row, col := round + 1, n - round - 1; row < n - round; row++ {
// 			res[row][col] = i
// 			i++
// 		}
// 		if i > total {
// 			break
// 		}
// 		// 下
// 		for row, col := n - round - 1, n - round - 2; col >= round; col-- {
// 			res[row][col] = i
// 			i++
// 		}
// 		if i > total {
// 			break
// 		}
// 		// 左
// 		for row, col := n - round - 2, round; row > round; row-- {
// 			res[row][col] = i
// 			i++
// 		}
// 		if i > total {
// 			break
// 		}
// 		round++
// 	}
// 	return res
// }

func generateMatrix(n int) [][]int {
	total := n * n
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	curDir := 0
	row, col := 0, 0
	for i := 0; i < total; i++ {
		res[row][col] = i + 1
		if r, c := row + dirs[curDir][0], col + dirs[curDir][1]; r >= n || r < 0 || c >= n || c < 0 || res[r][c] != 0 {
			curDir = (curDir + 1) % 4
		}
		row += dirs[curDir][0]
		col += dirs[curDir][1]
	}
	return res
}