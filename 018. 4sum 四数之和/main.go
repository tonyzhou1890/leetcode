package main

import (
	"leetcode/go-helper"
	"sort"
)

func main() {
	helper.PrintJSON(fourSum([]int{1,0,-1,0,-2,2}, 0)) // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	helper.PrintJSON(fourSum([]int{2,2,2,2,2}, 8)) // [[2,2,2,2]]
}

func fourSum(nums []int, target int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	res := [][]int{}
	length := len(nums)
	// 一层
	for i := 0; i < length - 3; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		if nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target {
			break
		}
		if nums[i] + nums[length - 3] + nums[length - 2] + nums[length - 1] < target {
			continue
		}
		// 二层
		for j := i + 1; j < length - 2; j++ {
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			if nums[i] + nums[j] + nums[j + 1] + nums[j + 2] > target {
				break
			}
			if nums[i] + nums[j] + nums[length - 2] + nums[length - 1] < target {
				continue
			}
			// 三层
			for k := j + 1; k < length - 1; k++ {
				if k > j + 1 && nums[k] == nums[k - 1] {
					continue
				}
				rest := target - nums[i] - nums[j] - nums[k]
				// 四层
				for l := k + 1; l < length; l++ {
					if nums[l] == rest {
						res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
						break
					}
					if nums[l] > rest {
						break
					}
				}
			}
		}
	}
	return res
}