package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(uniquePathsWithObstacles([][]int{{0,0,0},{0,1,0},{0,0,0}})) // 1
	helper.PrintJSON(uniquePathsWithObstacles([][]int{{0,1},{0,0}})) // 1
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if obstacleGrid[row][col] == 1 {
				continue
			}
			dp[row][col] = dp[row - 1][col] + dp[row][col - 1]
		}
	}
	return dp[m - 1][n - 1]
}