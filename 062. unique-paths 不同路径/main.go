package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(uniquePaths(3, 7)) // 28
	helper.PrintJSON(uniquePaths(3, 2)) // 3
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for col := 0; col < n; col++ {
		dp[0][col] = 1
	}
	for row := 0; row < m; row++ {
		dp[row][0] = 1
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = dp[row - 1][col] + dp[row][col - 1]
		}
	}
	return dp[m - 1][n - 1]
}