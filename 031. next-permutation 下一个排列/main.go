package main

import (
	"leetcode/go-helper"
)

func main() {
	temp := []int{1,2,3}
	nextPermutation(temp)
	helper.PrintJSON(temp) // [1,3,2]
	temp = []int{3,2,1}
	nextPermutation(temp)
	helper.PrintJSON(temp) // [1,2,3]
	temp = []int{1,1,5}
	nextPermutation(temp)
	helper.PrintJSON(temp) // [1,5,1]
	temp = []int{1,3,2}
	nextPermutation(temp)
	helper.PrintJSON(temp) // [2,1,3]
}

/**
最小值，递增；最大值，递减。从后面找到第一个局部峰值，峰值左侧即为需要交换的一个数 A。
A 后面从后往前查找，找到第一个比 A 大的数 B。交换 A 和 B
如果已经是最大值，逆转数组
 */
func nextPermutation(nums []int)  {
	length := len(nums)
	left := -1
	right := 0
	for i := length - 1; i > 0; i-- {
		// 峰值
		if nums[i] > nums[i - 1] {
			// 峰值左侧 A
			left = i - 1
			// A 右侧查找 B，使得 B > A
			for j := length - 1; j > left; j-- {
				if nums[j] > nums[left] {
					right = j
					break
				}
			}
			nums[left], nums[right] = nums[right], nums[left]
			break
		}
	}
	for i, j := left + 1, length - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}