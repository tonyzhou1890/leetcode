package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(twoSum([]int{2,7,11,15}, 9)) // [0,1]
	helper.PrintJSON(twoSum([]int{3,2,4}, 6)) // [1,2]
	helper.PrintJSON(twoSum([]int{3,3}, 6)) // [0,1]
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for idx, val := range nums {
		index, ok := hash[target - val]
		if ok {
			return []int{index, idx}
		}
		hash[val] = idx
	}
	return []int{}
}
