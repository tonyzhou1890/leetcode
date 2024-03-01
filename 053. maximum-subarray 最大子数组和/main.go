package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})) // 6
	helper.PrintJSON(maxSubArray([]int{1})) // 1
	helper.PrintJSON(maxSubArray([]int{5,4,-1,7,8})) // 23
}

func maxSubArray(nums []int) int {
	max := -(2 << 31)
	sum := 0
	for _, num := range nums {
		sum += num
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}