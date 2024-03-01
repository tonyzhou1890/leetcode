package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(removeElement([]int{3,2,2,3}, 3)) // 2
	helper.PrintJSON(removeElement([]int{0,1,2,2,3,0,4,2}, 2)) // 5
}

func removeElement(nums []int, val int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[p] = nums[i]
			p++
		}
	}
	return p
}