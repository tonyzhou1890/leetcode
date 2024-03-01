package main

import (
	"leetcode/go-helper"
	"sort"
)

func main() {
	helper.PrintJSON(permuteUnique([]int{1,1,2})) // [[1,1,2],[1,2,1],[2,1,1]]
	helper.PrintJSON(permuteUnique([]int{1,2,3})) // [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
}

// 通过字典序生成排列
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	sort.Sort(sort.IntSlice(nums))
	res = append(res, append([]int{}, nums...))
	for ;; {
		nums, isFirst := next(nums)
		if !isFirst {
			res = append(res, append([]int{}, nums...))
		} else {
			break
		}
	}
	return res
}

// 下一个排列。返回值：下一个排列，是否回到第一个排列
func next(nums []int) ([]int, bool) {
	l := -1
	r := 0
	isFirst := false
	// 找峰值左侧索引
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i - 1] {
			l = i - 1
			// 找峰值（包含）右侧最后一个比 l 处大的值
			for j := len(nums) - 1; j > l; j-- {
				if nums[j] > nums[l] {
					r = j
					break
				}
			}
			nums[l], nums[r] = nums[r], nums[l]
			break
		}
	}
	// 峰值右侧变升序
	for i, j := l + 1, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	if l == -1 {
		isFirst = true
	}
	return nums, isFirst
}