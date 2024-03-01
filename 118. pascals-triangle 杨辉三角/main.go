package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(generate(1))
	helper.PrintJSON(generate(2))
	helper.PrintJSON(generate(3))
	helper.PrintJSON(generate(4))
}


func generate(numRows int) [][]int {
	res := [][]int{}
	res = append(res, []int{1})

	for i := 1; i < numRows; i++ {
		lastRow := res[i - 1]
		currRow := make([]int, i + 1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				currRow[j] = 1
			} else {
				currRow[j] = lastRow[j] + lastRow[j - 1]
			}
		}
		res = append(res, currRow)
	}

	return res
}