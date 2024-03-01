package main

import (
	"leetcode/go-helper"
	"sort"
)

func main() {
	helper.PrintJSON(combinationSum2([]int{10,1,2,7,6,1,5}, 8)) 
	/*
	[
	[1,1,6],
	[1,2,5],
	[1,7],
	[2,6]
	]
	*/
	helper.PrintJSON(combinationSum2([]int{2,5,2,1,2}, 5))
	/*
	[
	[1,2,2],
	[5]
	]
	*/
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	sort.Sort(sort.IntSlice(candidates))
	length := len(candidates)

	var dfs func(candidates []int, start int, target int)
	dfs = func (candidates []int, start int, target int) {
		// 待选项不足
		if start == length {
			return
		}
		// 使用当前数字
		path = append(path, candidates[start])
		if target == candidates[start] {
			res = append(res, append([]int{}, path...))
		} else if target > candidates[start] {
			dfs(candidates, start + 1, target - candidates[start])
		}
		path = path[0 : len(path) - 1]
		// 不使用当前数字--后面相同数字全部跳过，因为这一轮使用当前数字的下一轮使用和不使用相同数字就包含了当前轮不使用当前数字的等价情况。
		offset := 1
		for i := start + 1; i < length; i++ {
			if candidates[i] == candidates[start] {
				offset++
			} else {
				break
			}
		}
		dfs(candidates, start + offset, target)
	}

	dfs(candidates, 0, target)
	return res
}