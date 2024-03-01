package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(subsets([]int{1,2,3})) // [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
	helper.PrintJSON(subsets([]int{0})) // [[],[0]]
}

// func subsets(nums []int) [][]int {
// 	res := [][]int{}
// 	sets := [][]int{}
// 	path := []int{}
// 	length := len(nums)

// 	help := func(int, int) {}
// 	help = func(start int, k int) {
// 		if k == 0 {
// 			sets = append(sets, append([]int{}, path...))
// 		} else {
// 			for ; start < length - k + 1; start++ {
// 				path = append(path, nums[start])
// 				help(start + 1, k - 1)
// 				path = path[ : len(path) - 1]
// 			}
// 		}
// 	}

// 	for i := 0; i <= length; i++ {
// 		sets = [][]int{}
// 		help(0, i)
// 		res = append(res, sets...)
// 	}

// 	return res
// }

// n 个数的子集可以是 n - 1 个数的子集全部加上新增加的数得到
func subsets(nums []int) [][]int {
	res := [][]int{}

	res = append(res, []int{})
	
	for i := 0; i < len(nums); i++ {
		for j, prevLen := 0, len(res); j < prevLen; j++ {
			res = append(res, append(append([]int{}, res[j]...), nums[i]))
		}
	}

	return res
}