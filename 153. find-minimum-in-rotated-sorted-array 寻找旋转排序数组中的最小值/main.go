package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(findMin([]int{3,4,5,1,2})) // 1
	helper.PrintJSON(findMin([]int{4,5,6,7,0,1,2})) // 0
	helper.PrintJSON(findMin([]int{11,13,15,17})) // 11
}

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	if nums[start] <= nums[end] {
		return nums[start]
	}

	idx := 0

	for start < end {
		mid := (start + end) / 2
		// 正好分割在突变处
		if nums[start] <= nums[mid] && nums[mid + 1] <= nums[end] {
			idx = mid + 1
			break
		}
		// 分割在突变左边
		if nums[start] < nums[mid] {
			start = mid
		} else { // 分割在突变右边
			end = mid
		}
	}

	return nums[idx]
}