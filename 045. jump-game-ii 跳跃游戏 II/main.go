package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(jump([]int{2,3,1,1,4})) // 2
	helper.PrintJSON(jump([]int{2,3,0,1,4})) // 2
	helper.PrintJSON(jump([]int{2})) // 0
}

// 每次的 end 作为一段的结束。到每段结束或者数组结束的时候，更新下一段的最远位置。
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	// 当前可到达的最远位置
	end := 0
	newEnd := 0
	step := 0
	for idx, val := range nums {
		if idx + val > newEnd {
			newEnd = idx + val
		}
		if end == idx {
			end = newEnd
			newEnd = 0
			step++
			if end >= len(nums) - 1 {
				break
			}
		}
	}
	return step
}