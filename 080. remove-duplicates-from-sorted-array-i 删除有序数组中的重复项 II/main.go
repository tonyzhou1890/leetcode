package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(removeDuplicates([]int{1,1,1,2,2,3})) // 5
	helper.PrintJSON(removeDuplicates([]int{0,0,1,1,1,1,2,3,3})) // 7
}

// func removeDuplicates(nums []int) int {
// 	length := len(nums)
// 	if length < 3 {
// 		return length
// 	}
// 	p1 := 1
// 	p2 := 2
// 	for ;p2 < length; p2++ {
// 		if nums[p1] == nums[p2] && nums[p1 - 1] == nums[p2] {
// 			continue
// 		}
// 		p1++
// 		nums[p1] = nums[p2]
// 	}
// 	return p1 + 1
// }

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 3 {
		return length
	}
	p := 2
	for i := 2; i < length; i++ {
		nums[p] = nums[i]
		// 填充数据有效，前进一步
		if nums[p - 2] != nums[p] {
			p++
		}
	}
	return p
}