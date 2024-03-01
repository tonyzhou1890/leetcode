package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(plusOne([]int{1,2,3})) // [1,2,4]
	helper.PrintJSON(plusOne([]int{4,3,2,1})) // [4,3,2,2]
	helper.PrintJSON(plusOne([]int{0})) // [1]
	helper.PrintJSON(plusOne([]int{9})) // [1,0]
}

func plusOne(digits []int) []int {
	carry := 1
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
		if carry == 0 {
			break
		}
	}
	if carry > 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}