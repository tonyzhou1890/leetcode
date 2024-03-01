package main

import (
	"leetcode/go-helper"
	"sort"
)

func main() {
	helper.PrintJSON(subsetsWithDup([]int{1,2,2})) // [[],[1],[1,2],[1,2,2],[2],[2,2]]
	helper.PrintJSON(subsetsWithDup([]int{0})) // [[],[0]]
	helper.PrintJSON(subsetsWithDup([]int{5,5,5,5,5})) // [[],[5],[5,5],[5,5,5],[5,5,5,5],[5,5,5,5,5]]
}

// 如果当前数字与上一个相同，且当前路径没有选择上一个数字，则跳过当前数字，因为第一次碰到这个数字的时候已经产生过对应路径了
func subsetsWithDup(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))

	res := [][]int{}

	path := []int{}

	var dfs func(start int, prevUsed bool)

	dfs = func (start int, prevUsed bool) {
		if start == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}

		dfs(start + 1, false)

		if start > 0 && nums[start - 1] == nums[start] && prevUsed == false {
			return
		}

		path = append(path, nums[start])

		dfs(start + 1, true)

		path = path[ : len(path) - 1]
	}

	dfs(0, false)

	return res
}