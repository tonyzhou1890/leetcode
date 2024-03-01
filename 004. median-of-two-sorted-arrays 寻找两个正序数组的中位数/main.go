package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(findMedianSortedArrays([]int{1, 3}, []int{2})) // 2.0000
	helper.PrintJSON(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5000
	helper.PrintJSON(findMedianSortedArrays([]int{1, 3, 4, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 4.5000
}

/**
 * 合并后的数组长度 m + n。中位数就是 (m + n) / 2。
 **/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	newLength := length1 + length2
	right := newLength / 2
	left := right
	mid := 0
	if newLength % 2 == 0 {
		left--
	}
	numsIdx1 := 0
	numsIdx2 := 0
	for i := 0; i <= right; i++ {
		// nums1 和 nums2 都没有走到尽头，比较当前的两个值
		if numsIdx1 < length1 && numsIdx2 < length2 {
			if nums1[numsIdx1] < nums2[numsIdx2] {
				if numsIdx1 + numsIdx2 == left || numsIdx1 + numsIdx2 == right {
					mid += nums1[numsIdx1]
				}
				numsIdx1++
			} else {
				if numsIdx1 + numsIdx2 == left || numsIdx1 + numsIdx2 == right {
					mid += nums2[numsIdx2]
				}
				numsIdx2++
			}
		} else if numsIdx1 < length1 {
			if numsIdx1 + numsIdx2 == left || numsIdx1 + numsIdx2 == right {
				mid += nums1[numsIdx1]
			}
			numsIdx1++
		} else {
			if numsIdx1 + numsIdx2 == left || numsIdx1 + numsIdx2 == right {
				mid += nums2[numsIdx2]
			}
			numsIdx2++
		}
	}
	if left == right {
		return float64(mid)
	} else {
		return (float64(mid) / 2)
	}
}