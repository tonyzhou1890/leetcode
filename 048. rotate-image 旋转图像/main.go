package main

import (
	"leetcode/go-helper"
)

func main() {
	arr := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	rotate(arr)
	helper.PrintJSON(arr) // [[7,4,1],[8,5,2],[9,6,3]]
	arr = [][]int{
		{5,1,9,11},
		{2,4,8,10},
		{13,3,6,7},
		{15,14,12,16},
	}
	rotate(arr)
	helper.PrintJSON(arr) // [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
}

func rotate(matrix [][]int)  {
	n := len(matrix)
	// 先沿着对角线（左下、右上）翻转
	for row := 0; row < n; row++ {
		for col := 0; col < n - row; col++ {
			matrix[n - col - 1][n - row - 1], matrix[row][col] = matrix[row][col], matrix[n - col - 1][n - row - 1]
		}
	}
	// 再上下反转
	for row := 0; row < n / 2; row++ {
		for col := 0; col < n; col++ {
			matrix[n - row - 1][col], matrix[row][col] = matrix[row][col], matrix[n - row - 1][col]
		}
	}
}