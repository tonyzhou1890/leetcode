package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(singleNumber([]int{2,2,1})) // 1
	helper.PrintJSON(singleNumber([]int{4,1,2,1,2})) // 4
	helper.PrintJSON(singleNumber([]int{1})) // 1
}

// 异或法
// a ^ a = 0; a ^ b ^ a = b
func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}