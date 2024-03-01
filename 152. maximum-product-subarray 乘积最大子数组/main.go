package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(maxProduct([]int{2,3,-2,4})) // 6
	helper.PrintJSON(maxProduct([]int{-2,0,-1})) // 0
	helper.PrintJSON(maxProduct([]int{3,-1,4})) // 4
	helper.PrintJSON(maxProduct([]int{-2,-3,7})) // 42
	helper.PrintJSON(maxProduct([]int{-4,-3,-2})) // 12
}

func maxProduct(nums []int) int {
	min := 1
	max := 1
	res := -2 << 31
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		maxTemp := max
		if min * nums[i] > temp {
			temp = min * nums[i]
		}
		if max * nums[i] > temp {
			temp = max * nums[i]
		}
		max = temp

		temp = nums[i]
		if min * nums[i] < temp {
			temp = min * nums[i]
		}
		if maxTemp * nums[i] < temp {
			temp = maxTemp * nums[i]
		}
		min = temp

		if max > res {
			res = max
		}
	}
	return res
}