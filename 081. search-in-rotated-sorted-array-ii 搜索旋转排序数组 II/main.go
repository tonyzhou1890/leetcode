package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(search([]int{2,5,6,0,0,1,2}, 0)) // true
	helper.PrintJSON(search([]int{2,5,6,0,0,1,2}, 3)) // false
	helper.PrintJSON(search([]int{1,0,1,1,1}, 0)) // true
	helper.PrintJSON(search([]int{3,1}, 1)) // true
}

func search(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		} else if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
		} else if nums[mid] > target {
			// 左侧有序
			if nums[l] <= nums[mid] {
				if nums[l] <= target {
					r = mid - 1
				} else {
					l = mid + 1
				}
			} else {
				// 右侧有序
				r = mid - 1
			}
		} else {
			// 左侧有序
			if nums[l] <= nums[mid] {
				l = mid + 1
			} else {
				// 右侧有序
				if nums[r] >= target {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
		}
	}
	return false
}