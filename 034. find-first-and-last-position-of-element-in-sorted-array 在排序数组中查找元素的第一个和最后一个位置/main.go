package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(searchRange([]int{5,7,7,8,8,10}, 8)) // [3, 4]
	helper.PrintJSON(searchRange([]int{5,7,7,8,8,10}, 6)) // [-1, -1]
	helper.PrintJSON(searchRange([]int{}, 0)) // [-1. -1]
	helper.PrintJSON(searchRange([]int{1,2,3}, 2)) // [1, 1]
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	l := 0
	r := len(nums) - 1
	// 左边界
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] != target {
		return res
	} else {
		res[0] = l
	}
	l = 0
	r = len(nums) - 1
	// 有边界
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		res[1] = l
	} else if nums[l - 1] == target {
		res[1] = l - 1
	}
	return res
}