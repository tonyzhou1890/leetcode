package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(minPathSum([][]int{{1,3,1},{1,5,1},{4,2,1}})) // 7
	helper.PrintJSON(minPathSum([][]int{{1,2,3},{4,5,6}})) // 12
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for col := 1; col < n; col++ {
		grid[0][col] += grid[0][col - 1]
	}
	for row := 1; row < m; row++ {
		grid[row][0] += grid[row - 1][0]
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if grid[row - 1][col] < grid[row][col - 1] {
				grid[row][col] += grid[row - 1][col]
			} else {
				grid[row][col] += grid[row][col - 1]
			}
		}
	}

	return grid[m - 1][n - 1]
}