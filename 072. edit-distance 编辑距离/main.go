package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(minDistance("horse", "ros")) // 3
	helper.PrintJSON(minDistance("intention", "execution")) // 5
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	// dp[i][j] 代表 word1 截止到 i 位置变换到 word2 截止到 j 位置的编辑距离
	// 多的一行一列代表空字符串
	dp := make([][]int, m + 1)
	for row := 0; row <= m; row++ {
		dp[row] = make([]int, n + 1)
	}
	for col := 0; col <= n; col++ {
		dp[0][col] = col
	}
	for row := 0; row <= m; row++ {
		dp[row][0] = row
	}
	for row := 1; row <= m; row++ {
		for col := 1; col <= n; col++ {
			// 字符相同，则从 dp[row - 1][col - 1] 到 dp[row][col] 不需要变换
			if word1[row - 1] == word2[col - 1] {
				dp[row][col] = getMin(dp[row - 1][col] + 1, dp[row][col - 1] + 1, dp[row - 1][col - 1])
			} else {
				dp[row][col] = getMin(dp[row - 1][col] + 1, dp[row][col - 1] + 1, dp[row - 1][col - 1] + 1)
			}
		}
	}
	return dp[m][n]
}

func getMin(a int, b int, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}