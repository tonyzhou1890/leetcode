package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 3)) // true
	helper.PrintJSON(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 13)) // false
}

// 二分搜索
// func searchMatrix(matrix [][]int, target int) bool {
// 	m := len(matrix)
// 	n := len(matrix[0])
// 	l := 0
// 	r := m * n - 1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		row, col := toXY(m, n, mid)
// 		if matrix[row][col] == target {
// 			return true
// 		}
// 		if matrix[row][col] < target {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return false
// }

// func toXY(m int, n int, idx int) (int, int) {
// 	return idx / n, idx % n
// }

// 从第一行的最后开始查找，目标值可能在这一行，也可能在这一行下面。等于当前行最后一个元素，直接返回。小于当前行最后一个元素，说明在当前行。否则在当前行下面。
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	row := 0
	col := n - 1
	for col >= 0 && row < m {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	return false
}
