package main

import (
	"leetcode/go-helper"
)

func main() {
	// arr := []int{2,2,2,2,2,2};
	arr := [][]int{
		{0,1,0},
		{0,0,1},
		{1,1,1},
		{0,0,0}};
	gameOfLife(arr)
	helper.PrintJSON(arr)
}

/**
九宫格内存在 3-4 个活细胞，并且中间的细胞，则为活细胞，否则为死细胞。如果中间为死细胞，九宫格内正好存在三个活细胞，则为活细胞，否则死细胞。
用 int 数值的倒数第二位存储计算后的结果。
*/
func gameOfLife(board [][]int)  {
	yLen := len(board)
	if yLen == 0 {
		return
	}
	xLen := len(board[0])
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			count := 0
			xMin := x - 1
			if xMin < 0 {
				xMin = 0
			}
			xMax := x + 1
			if xMax >= xLen {
				xMax = xLen - 1
			}
			yMin := y - 1
			if yMin < 0 {
				yMin = 0
			}
			yMax := y + 1
			if yMax >= yLen {
				yMax = yLen - 1
			}
			for yIdx := yMin; yIdx <= yMax; yIdx++ {
				for xIdx := xMin; xIdx <= xMax; xIdx++ {
					if (board[yIdx][xIdx] & 0b1) == 1 {
						count++
					}
				}
			}
			if (board[y][x] & 0b1) == 1 {
				if (count == 3 || count == 4) {
					board[y][x] = board[y][x] | 0b10
				}
			} else if count == 3 {
				board[y][x] = board[y][x] | 0b10
			}
		}
	}
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			board[y][x] = board[y][x] >> 1
		}
	}
}
