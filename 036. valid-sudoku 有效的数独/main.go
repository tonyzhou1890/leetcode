package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(isValidSudoku([][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'}})) // true
	helper.PrintJSON(isValidSudoku([][]byte{
		{'8','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'}})) // false
}

// func isValidSudoku(board [][]byte) bool {
// 	rowMap := [9]map[byte]bool{}
// 	colMap := [9]map[byte]bool{}
// 	cellMap := [9]map[byte]bool{}
// 	for row := 0; row < 9; row++ {
// 		for col := 0; col < 9; col++ {
// 			cur := board[row][col]
// 			if cur != '.' {
// 				// row
// 				if _, ok := rowMap[row][cur]; ok {
// 					return false
// 				} else {
// 					if rowMap[row] == nil {
// 						rowMap[row] = map[byte]bool{}
// 					}
// 					rowMap[row][cur] = true
// 				}
// 				// col 
// 				if _, ok := colMap[col][cur]; ok {
// 					return false
// 				} else {
// 					if colMap[col] == nil {
// 						colMap[col] = map[byte]bool{}
// 					}
// 					colMap[col][cur] = true
// 				}
// 				// cell
// 				cellIdx := row / 3 * 3 + col / 3
// 				if _, ok := cellMap[cellIdx][cur]; ok {
// 					return false
// 				} else {
// 					if cellMap[cellIdx] == nil {
// 						cellMap[cellIdx] = map[byte]bool{}
// 					}
// 					cellMap[cellIdx][cur] = true
// 				}
// 			}
// 		}
// 	}
// 	return true
// }

func isValidSudoku(board [][]byte) bool {
	rowMap := [9][9]bool{}
	colMap := [9][9]bool{}
	cellMap := [3][3][9]bool{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				cur := board[row][col] - '1'
				if rowMap[row][cur] == true || colMap[col][cur] == true || cellMap[row / 3][col / 3][cur] == true {
					return false
				}
				rowMap[row][cur] = true
				colMap[col][cur] = true
				cellMap[row / 3][col / 3][cur] = true
			}
		}
	}
	return true
}