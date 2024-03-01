package main

import (
	"leetcode/go-helper"
	"math"
	"sort"
)

func main() {
	helper.PrintJSON(threeSumClosest([]int{-1,2,1,-4}, 1)) // 2
	helper.PrintJSON(threeSumClosest([]int{0,0,0}, 1)) // 0
}

// threeSumClosest
func threeSumClosest(nums []int, target int) int {
	res := 2 << 31
	sort.Sort(sort.IntSlice(nums))
	length := len(nums)
	for i := 0; i < length - 2; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		l, r := i + 1, length - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			} else if sum > target {
				res = closet(sum, res, target)
				r--
			} else {
				res = closet(sum, res, target)
				l++
			}
		}

	}

	return res
}

func closet(curr int, prev int, target int) int {
	if math.Abs(float64(curr) - float64(target)) < math.Abs(float64(prev) - float64(target)) {
		return curr
	}
	return prev
}