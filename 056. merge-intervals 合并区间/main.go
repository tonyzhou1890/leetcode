package main

import (
	"leetcode/go-helper"
	"sort"
)

func main() {
	helper.PrintJSON(merge([][]int{{1,3},{2,6},{8,10},{15,18},})) // [[1,6],[8,10],[15,18]]
	helper.PrintJSON(merge([][]int{{1,4},{4,5},})) // [[1,5]]
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	res := [][]int{}
	cur := []int{}
	for i := 0; i < len(intervals); i++ {
		if len(cur) == 0 {
			cur = intervals[i]
			continue
		}
		if intervals[i][0] <= cur[1] {
			if intervals[i][1] > cur[1] {
				cur[1] = intervals[i][1]
			}
		} else {
			res = append(res, append([]int{}, cur...))
			cur = intervals[i]
		}
	}
	if cur != nil {
		res = append(res, cur)
	}
	return res
}