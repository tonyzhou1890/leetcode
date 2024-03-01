package main

import (
	"leetcode/go-helper"
)

func main() {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}
	merge(nums1, 3, nums2, 3)
	helper.PrintJSON(nums1) // [1,2,2,3,5,6]
	nums1 = []int{1}
	nums2 = []int{}
	merge(nums1, 1, nums2, 0)
	helper.PrintJSON(nums1) // [1]
	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	helper.PrintJSON(nums1) // [1]
}

// 两个数组都是升序，并且 nums1 空间足够。所以，从最后面插入数据。
func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i, j, k := m - 1, n - 1, m + n - 1; j >= 0; k-- {
		if i > 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}