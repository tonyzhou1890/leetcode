package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9},})) // [1,2,3,6,9,8,7,4,5]
	helper.PrintJSON(spiralOrder([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},})) // [1,2,3,4,8,12,11,10,9,5,6,7]
}

func spiralOrder(matrix [][]int) []int {
	left, top := 0, 0
	right, bottom := len(matrix[0]), len(matrix)
	res := []int{}
	for left < right && top < bottom {
		for row, col := top, left; col < right; col++ {
			res = append(res, matrix[row][col])
		}
		for row, col := top + 1, right - 1; row < bottom; row++ {
			res = append(res, matrix[row][col])
		}
		if top < bottom - 1 && left < right - 1 {
			for row, col := bottom - 1, right - 2; col >= left; col-- {
				res = append(res, matrix[row][col])
			}
			for row, col := bottom - 2, left; row > top; row-- {
				res = append(res, matrix[row][col])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}