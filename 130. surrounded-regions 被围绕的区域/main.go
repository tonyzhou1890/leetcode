package main

import (
	// "leetcode/go-helper"
)

func main() {
}

func solve(board [][]byte)  {
	maxRow := len(board)
	maxCol := len(board[0])
	// 把相连字符设为指定字符
	var setFlag func(row int, col int, rawC byte, targetC byte)
	setFlag = func(row, col int, rawC byte, targetC byte) {
		board[row][col] = targetC
		if row > 0 && board[row - 1][col] == rawC {
			setFlag(row - 1, col, rawC, targetC)
		}
		if col < maxCol - 1 && board[row][col + 1] == rawC {
			setFlag(row, col + 1, rawC, targetC)
		}
		if row < maxRow - 1 && board[row + 1][col] == rawC {
			setFlag(row + 1, col, rawC, targetC)
		}
		if col > 0 && board[row][col - 1] == rawC {
			setFlag(row, col - 1, rawC, targetC)
		}
	}
	// 把边界相连的 O 全部置为 B
	for col := 0; col < maxCol; col++ {
		if board[0][col] == 'O' {
			setFlag(0, col, 'O', 'B')
		}
		if board[maxRow - 1][col] == 'O' {
			setFlag(maxRow - 1, col, 'O', 'B')
		}
	}
	for row := 0; row < maxRow; row++ {
		if board[row][0] == 'O' {
			setFlag(row, 0, 'O', 'B')
		}
		if board[row][maxCol - 1] == 'O' {
			setFlag(row, maxCol - 1, 'O', 'B')
		}
	}
	// 把剩下的 O 置为 X，B 置为 O
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			} else if board[row][col] == 'B' {
				board[row][col] = 'O'
			}
		}
	}
}