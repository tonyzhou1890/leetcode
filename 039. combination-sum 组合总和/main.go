package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(combinationSum([]int{2,3,6,7}, 7)) // [[2,2,3],[7]]
	helper.PrintJSON(combinationSum([]int{2,3,5}, 8)) // [[2,2,2,2],[2,3,3],[3,5]]
	helper.PrintJSON(combinationSum([]int{2}, 1)) // []
	helper.PrintJSON(combinationSum([]int{8,7,4,3}, 11)) // [[8,3],[7,4],[4,4,3]]
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	// 待选项不足
	if len(candidates) == 0 {
		return res
	}
	// 递归
	// 使用当前数字
	if target == candidates[0] {
		res = append(res, []int{candidates[0]})
	} else if target > candidates[0] {
		useRes := combinationSum(candidates[0 : ], target - candidates[0])
		for _, val := range useRes {
			res = append(res, append([]int{candidates[0]}, val...))
		}
	}
	// 不使用当前数字
	notUseRes := combinationSum(candidates[1 : ], target)
	for _, val := range notUseRes {
		res = append(res, val)
	}
	return res
}