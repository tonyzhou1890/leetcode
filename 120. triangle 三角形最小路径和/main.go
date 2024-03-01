package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(minimumTotal([][]int{{2},{3,4},{6,5,7},{4,1,8,3}})) // 11
	helper.PrintJSON(minimumTotal([][]int{{-10}})) // -10
}


// func minimumTotal(triangle [][]int) int {
// 	dp := make([]int, len(triangle))
// 	for row := 0; row < len(triangle); row++ {
// 		// 每一行从最后一个开始，因为从第一个开始会直接覆盖第一个结果，导致第二个无法计算
// 		for col := row; col >= 0; col-- {
// 			if col == 0 {
// 				dp[col] += triangle[row][col]
// 			} else if col == row {
// 				dp[col] = triangle[row][col] + dp[col - 1]
// 			} else {
// 				if dp[col] < dp[col - 1] {
// 					dp[col] += triangle[row][col]
// 				} else {
// 					dp[col] = dp[col - 1] + triangle[row][col]
// 				}
// 			}
// 		}
// 	}
// 	min := dp[0]
// 	for i := 1; i < len(dp); i++ {
// 		if dp[i] < min {
// 			min = dp[i]
// 		}
// 	}

// 	return min
// }

// 自底向上，修改原数组
func minimumTotal(triangle [][]int) int {
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			if triangle[row + 1][col] < triangle[row + 1][col + 1] {
				triangle[row][col] += triangle[row + 1][col]
			} else {
				triangle[row][col] += triangle[row + 1][col + 1]
			}
		}
	}
	return triangle[0][0]
}