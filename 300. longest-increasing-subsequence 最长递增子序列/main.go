package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(lengthOfLIS(helper.GenerateRandomArray(2500, -10000, 10000)))
	helper.PrintJSON(lengthOfLIS([]int{10,9,2,5,3,7,101,18})) // 4
	helper.PrintJSON(lengthOfLIS([]int{0,1,0,3,2,3})) // 4
	helper.PrintJSON(lengthOfLIS([]int{7,7,7,7,7,7,7})) // 1
	helper.PrintJSON(lengthOfLIS([]int{4,10,4,3,8,9})) // 3
}

// func lengthOfLIS(nums []int) int {
// 	min := ^int(^uint(0) >> 1)
// 	cache := map[string]int{}
// 	return subArrCount(&nums, 0, min, len(nums), &cache)
// }

// func subArrCount(nums *[]int, startIdx int, compareNum int, length int, cache *map[string]int) int {
// 	// get cache
// 	key := strconv.Itoa(startIdx) + "-" + strconv.Itoa(compareNum)
// 	if val, ok := (*cache)[key]; ok {
// 		return val
// 	}
// 	if startIdx >= length {
// 		return 0
// 	}
// 	containCount := 0
// 	if (*nums)[startIdx] > compareNum {
// 		// 计数当前数字
// 		containCount = 1 + subArrCount(nums, startIdx + 1, (*nums)[startIdx], length, cache)
// 	}
// 	// 不计数当前数字
// 	notContainCount := subArrCount(nums, startIdx + 1, compareNum, length, cache)
// 	if containCount > notContainCount {
// 		(*cache)[key] = containCount
// 		return containCount
// 	}
// 	(*cache)[key] = notContainCount
// 	return notContainCount
// }

/**
动态规划
dp[i] 代表以 nums[i] 结尾的最长上升子序列长度
dp[i] = max(dp[j])。其中，0 <= j < i, 且 nums[j] < nums[i]
*/
// func lengthOfLIS(nums []int) int {
// 	dp := []int{}
// 	for key, value := range nums {
// 		maxLen := 0
// 		for j := 0; j < key; j++ {
// 			if nums[j] < value && dp[j] > maxLen {
// 				maxLen = dp[j]
// 			}
// 		}
// 		dp = append(dp, maxLen + 1)
// 	}
// 	res := 0
// 	for _, value := range dp {
// 		if value > res {
// 			res = value
// 		}
// 	}
// 	return res
// }

/**
贪心+二分查找
dp[i] 代表长度为 i 的子序列结尾最小的数
*/
func lengthOfLIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length + 1)
	if length == 0 {
		return 0
	}
	currLen := 1
	dp[currLen] = nums[0]
	for i := 1; i < length; i++ {
		if nums[i] > dp[currLen] {
			currLen++
			dp[currLen] = nums[i]
		} else {
			left := 1
			right := currLen
			pos := 0
			for left <= right {
				mid := (right - left) / 2 + left
				if dp[mid] < nums[i] {
					left = mid + 1
					pos = mid
				} else {
					right = mid - 1
				}
			}
			dp[pos + 1] = nums[i]
		}
	}
	return currLen
}
