package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "ABCCED")) // true
	helper.PrintJSON(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "SEE")) // true
	helper.PrintJSON(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "ABCB")) // false
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	sLen := len(word)

	tag := [][]bool{}
	for i := 0; i < m; i++ {
		tag = append(tag, make([]bool, n))
	}

	var dfs func(int, int, int) bool
	dfs = func(row int, col int, idx int) bool {
		if board[row][col] != word[idx] || tag[row][col] {
			return false
		}
		if board[row][col] == word[idx] {
			if idx == sLen - 1 {
				return true
			}
			tag[row][col] = true
			// 向右
			if col < n - 1 && dfs(row, col + 1, idx + 1) {
				return true
			}
			// 向下
			if row < m - 1 && dfs(row + 1, col, idx + 1) {
				return true
			}
			// 向左
			if col > 0 && dfs(row, col - 1, idx + 1) {
				return true
			}
			// 向上
			if row > 0 && dfs(row - 1, col, idx + 1) {
				return true
			}
			tag[row][col] = false
		}
		return false
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if dfs(row, col, 0) {
				return true
			}
		}
	}
	return false
}