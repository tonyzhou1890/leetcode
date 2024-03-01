package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(maxArea([]int{1,8,6,2,5,4,8,3,7})) // 49
	helper.PrintJSON(maxArea([]int{1,1})) // 1
}

func maxArea(height []int) int {
	max := 0
	area := 0
	i := 0
	j := len(height) - 1
	for i < j {
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if area > max {
			max = area
		}
	}
	return max
}
