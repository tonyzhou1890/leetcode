package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(insert([][]int{{1,3},{6,9},}, []int{2,5})) // [[1,5],[6,9]]
	helper.PrintJSON(insert([][]int{{1,2},{3,5},{6,7},{8,10},{12,16},}, []int{4,8})) // [[1,2],[3,10],[12,16]]
	helper.PrintJSON(insert([][]int{}, []int{5,7})) // [[5,7]]
	helper.PrintJSON(insert([][]int{{1,5}}, []int{2,3})) // [[1,5]]
	helper.PrintJSON(insert([][]int{{1,5}}, []int{2,7})) // [[1,7]]
}

// 先插入，在合并区间
func insert(intervals [][]int, newInterval []int) [][]int {
	insertIdx := -1
	for idx, val := range intervals {
		if val[0] > newInterval[0] {
			intervals = append(append(append([][]int{}, intervals[ : idx]...), newInterval), intervals[idx :]...)
			insertIdx = idx
			break
		}
	}
	if insertIdx == -1 {
		intervals = append(intervals, newInterval)
		insertIdx = len(intervals) - 1
	}
	startIdx := insertIdx
	// 是否与前一个区间重叠
	if insertIdx != 0 && intervals[insertIdx - 1][1] >= intervals[insertIdx][0] {
		startIdx = insertIdx - 1
	}
	endIdx := startIdx
	// 找到重叠区间结束位置
	for i := startIdx + 1; i < len(intervals); i++ {
		if intervals[startIdx][1] >= intervals[i][0] {
			if intervals[startIdx][1] < intervals[i][1] {
				intervals[startIdx][1] = intervals[i][1]
			}
			
			endIdx = i
		} else {
			break
		}
	}
	return append(intervals[ : startIdx + 1], intervals[endIdx + 1 :]...)
}