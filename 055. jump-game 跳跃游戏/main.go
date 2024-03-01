package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(canJump([]int{2,3,1,1,4})) // true
	helper.PrintJSON(canJump([]int{3,2,1,0,4})) // false
}

func canJump(nums []int) bool {
	// 当前可到达的最远位置
	end := 0
	for idx, val := range nums {
		// 无法达到当前位置，结束
		if end < idx {
			return false
		}
		if idx + val > end {
			end = idx + val
		}
	}
	return true
}