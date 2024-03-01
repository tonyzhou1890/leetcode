package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(searchInsert([]int{1,3,5,6}, 5)) // 2
	helper.PrintJSON(searchInsert([]int{1,3,5,6}, 2)) // 1
	helper.PrintJSON(searchInsert([]int{1,3,5,6}, 7)) // 4
}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}