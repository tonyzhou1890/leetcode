package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(combine(4, 2)) // [[2,4],[3,4],[2,3],[1,2],[1,3],[1,4],]
	helper.PrintJSON(combine(1, 1)) // [[1]]
}

func combine(n int, k int) [][]int {
	res := [][]int{}
	temp := [][]int{}
	for count := 0; count < k; count++ {
		temp = [][]int{}
		if count == 0 {
			for i := 0; i < n - (k - count) + 1; i++ {
				temp = append(temp, []int{i + 1})
			}
		} else {
			for _, val := range res {
				for i := val[count - 1]; i < n - (k - count) + 1; i++ {
					temp = append(temp, append(append([]int{}, val...), i + 1))
				}
			}
		}
		res = temp
	}
	return res
}
