package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(singleNumber([]int{2,2,3,2})) // 3
	helper.PrintJSON(singleNumber([]int{0,1,0,1,0,1,99})) // 99
	helper.PrintJSON(singleNumber([]int{-2,-2,1,1,4,1,4,4,-4,-2})) // -4
}

// 按位计数，除了出现一次的数，其余的数每位的个数和应该是 3 的倍数，算上出现一次的数，不是三的倍数的，则说明这一位存在于出现一次的数中
func singleNumber(nums []int) int {
	res := int32(0)
	count := [32]int{}
	for i := 0; i < 32; i++ {
		for j := 0; j < len(nums); j++ {
			count[i] += (nums[j] >> (31 - i)) & 1
		}
		res = (res << 1) + int32(count[i] % 3)
	}
	return int(res)
}