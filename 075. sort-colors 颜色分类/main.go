package main

import (
	"leetcode/go-helper"
)

func main() {
	arr := []int{2,0,2,1,1,0}
	sortColors(arr)
	helper.PrintJSON(arr) // [0,0,1,1,2,2]
	arr = []int{2,0,1}
	sortColors(arr)
	helper.PrintJSON(arr) // [0,1,2]
}

// // 快排
// func sortColors(nums []int)  {
// 	length := len(nums)
// 	if length <= 1 {
// 		return
// 	}

// 	quickSort(&nums, 0, length - 1)
// }

// func quickSort(nums *[]int, start int, end int) {
// 	if start >= end {
// 		return
// 	}
// 	pivotIdx := partition(nums, start, end)
// 	quickSort(nums, start, pivotIdx - 1)
// 	quickSort(nums, pivotIdx + 1, end)
// }

// func partition(nums *[]int, start int, end int) int {
// 	pivot := (*nums)[start]
// 	l := start
// 	r := end
// 	for l < r {
// 		for l < r && (*nums)[r] > pivot {
// 			r--
// 		}
// 		for l < r && (*nums)[l] <= pivot {
// 			l++
// 		}
// 		(*nums)[l], (*nums)[r] = (*nums)[r], (*nums)[l]
// 	}

// 	(*nums)[l], (*nums)[start] = (*nums)[start], (*nums)[l]
// 	return l
// }

// 堆排序
func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	heapify(&nums, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(&nums, 0, i)
	}
}

func heapify(nums *[]int, node int, length int) {
	l := (node << 1) + 1
	r := l + 1
	largest := node
	if l < length {
		heapify(nums, l, length)
		if (*nums)[l] > (*nums)[largest] {
			largest = l
		}
	}
	if r < length {
		heapify(nums, r, length)
		if (*nums)[r] > (*nums)[largest] {
			largest = r
		}
	}
	(*nums)[node], (*nums)[largest] = (*nums)[largest], (*nums)[node]
}