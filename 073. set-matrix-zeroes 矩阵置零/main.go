package main

import (
	"leetcode/go-helper"
)

func main() {
	arr := [][]int{{1,1,1},{1,0,1},{1,1,1}}
	setZeroes(arr)
	helper.PrintJSON(arr) // [[1,0,1],[0,0,0],[1,0,1]]
	arr = [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}
	setZeroes(arr)
	helper.PrintJSON(arr) // [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
}

/*
用两个数组记录哪些行和列需要置0。
换个想法，用原数组的第一行和第一列记录哪些行和列需要置0。
但这就需要提前记录第一行和第一列是否有原始的0，之后置零的时候需要知道第一行和第一列是否需要置0.
*/
// func setZeroes(matrix [][]int)  {
// 	m := len(matrix)
// 	n := len(matrix[0])
// 	rowFlag := false
// 	colFlag := false
// 	// 记录第一行是否有原始 0
// 	for col := 0; col < n; col++ {
// 		if matrix[0][col] == 0 {
// 			rowFlag = true
// 			break
// 		}
// 	}
// 	// 记录第一列是否有原始 0
// 	for row := 0; row < m; row++ {
// 		if matrix[row][0] == 0 {
// 			colFlag = true
// 			break
// 		}
// 	}
// 	// 从 1,1 开始遍历，记录行列是否需要置 0
// 	for row := 1; row < m; row++ {
// 		for col := 1; col < n; col++ {
// 			if matrix[row][col] == 0 {
// 				matrix[row][0] = 0
// 				matrix[0][col] = 0
// 			}
// 		}
// 	}
// 	// 行列置 0
// 	for row := 1; row < m; row++ {
// 		if matrix[row][0] == 0 {
// 			for col := 1; col < n; col++ {
// 				matrix[row][col] = 0
// 			}
// 		}
// 	}
// 	for col := 1; col < n; col++ {
// 		if matrix[0][col] == 0 {
// 			for row := 1; row < m; row++ {
// 				matrix[row][col] = 0
// 			}
// 		}
// 	}
// 	// 第一行和第一列置 0
// 	if rowFlag {
// 		for col := 0; col < n; col++ {
// 			matrix[0][col] = 0
// 		}
// 	}
// 	if colFlag {
// 		for row := 0; row < m; row++ {
// 			matrix[row][0] = 0
// 		}
// 	}
// }

func setZeroes(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])
	rowFlag := false
	colFlag := false
	// 记录第一行是否有原始 0
	for col := 0; col < n; col++ {
		if matrix[0][col] == 0 {
			rowFlag = true
			break
		}
	}
	// 记录第一列是否有原始 0
	for row := 0; row < m; row++ {
		if matrix[row][0] == 0 {
			colFlag = true
			break
		}
	}
	// 从 1,1 开始遍历，记录行列是否需要置 0
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}
	// 行列置 0
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
	}
	// 第一行和第一列置 0
	if rowFlag {
		for col := 0; col < n; col++ {
			matrix[0][col] = 0
		}
	}
	if colFlag {
		for row := 0; row < m; row++ {
			matrix[row][0] = 0
		}
	}
}