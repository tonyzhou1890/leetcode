package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(search([]int{4,5,6,7,0,1,2}, 0)) // 4
	helper.PrintJSON(search([]int{4,5,6,7,0,1,2}, 3)) // -1
	helper.PrintJSON(search([]int{3}, 0)) // -1
}

/**
从中间分为两部分，一部分递增，一部分包含旋转点。递增的部分二分，包含旋转点的递归。
 */
func search(nums []int, target int) int {
	length := len(nums)
	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	start := 0
	end := length - 1
	mid := (end - start) / 2 + start
	leftRes := -1
	rightRes := -1
	if nums[start] < nums[mid] {
		leftRes = bSearch(nums[start : mid + 1], target)
		rightRes = search(nums[mid + 1 : ], target)
	} else {
		leftRes = search(nums[start : mid + 1], target)
		rightRes = bSearch(nums[mid + 1 : ], target)
	}
	if leftRes != -1 {
		return leftRes
	}
	if rightRes != -1 {
		return mid + 1 + rightRes
	}
	return -1
}

func bSearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := (end - start) / 2 + start
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if nums[start] == target {
		return start
	}
	return -1
}