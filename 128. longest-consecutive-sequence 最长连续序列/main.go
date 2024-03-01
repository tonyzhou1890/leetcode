package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(longestConsecutive([]int{100,4,200,1,3,2})) // 4
	helper.PrintJSON(longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1})) // 9
}

/**
哈希表
第一轮建立哈希表
第二轮判断以 x 为起点的最长长度（如果存在 x - 1，则放弃本轮，因为 x - 1 开始的连续长度肯定比 x 开始的长）
*/
// func longestConsecutive(nums []int) int {
// 	hashTable := map[int]bool{}
// 	maxLength := 0

// 	for _, val := range nums {
// 		hashTable[val] = true
// 	}

// 	for _, val := range nums {
// 		// 存在更小的数，跳过
// 		if _, ok := hashTable[val - 1]; ok {
// 			continue
// 		}
// 		// 计数
// 		max := val + 1
// 		for true {
// 			if _, ok := hashTable[max]; !ok {
// 				break
// 			}
// 			max++
// 		}
// 		if max - val > maxLength {
// 			maxLength = max - val
// 		}
// 	}

// 	return maxLength
// }

func longestConsecutive(nums []int) int {
	hashTable := make(map[int]bool, len(nums))
	maxLength := 0

	for i := 0; i < len(nums); i++ {
		hashTable[nums[i]] = true
	}

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		// 存在更小的数，跳过
		if hashTable[val - 1] {
			continue
		}
		// 计数
		max := val + 1
		for true {
			if hashTable[max] {
				max++
			} else {
				break
			}
		}
		if max - val > maxLength {
			maxLength = max - val
		}
	}

	return maxLength
}