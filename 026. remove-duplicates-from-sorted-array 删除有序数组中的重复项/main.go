package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(removeDuplicates([]int{1,1,2})) // 2
	helper.PrintJSON(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})) // 5
}

func removeDuplicates(nums []int) int {
	p := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[p] {
			nums[p + 1] = nums[i]
			p++
		}
	}
	return p + 1
}