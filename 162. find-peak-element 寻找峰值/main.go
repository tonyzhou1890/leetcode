package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(findPeakElement([]int{1,2,3,1})) // 2
	helper.PrintJSON(findPeakElement([]int{1,2,1,3,5,6,4})) // 1 , 5
}

// func findPeakElement(nums []int) int {
// 	for i := 0; i < len(nums) - 1; i++ {
// 		if nums[i] > nums[i + 1] {
// 			return i
// 		}
// 	}
// 	return len(nums) - 1
// }

/**
二分
因为 nums[-1] 和 nums[nums.length] 为 -∞，所以 nums 里一定存在峰值。
如果 nums[mid] < nums[mid + 1]，说明右侧上坡，走右侧
如果 nums[mid] < nums[mid - 1]，说明左侧下坡，走左侧
否则就是中点
 */
func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[mid + 1] {
			l = mid + 1
		} else if mid > 0 && nums[mid] < nums[mid - 1] {
			r = mid
		} else {
			return mid
		}
	}
	return l
}
