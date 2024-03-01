package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(maxProfit([]int{7,1,5,3,6,4})) // 7
	helper.PrintJSON(maxProfit([]int{1,2,3,4,5})) // 4
	helper.PrintJSON(maxProfit([]int{7,6,4,3,1})) // 0
}


func maxProfit(prices []int) int {
	min := prices[0]
	max := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		// 开始新的区间段
		if prices[i] < max {
			profit += max - min
			min, max = prices[i], prices[i]
		} else if prices[i] > max {
			max = prices[i]
		}
	}
	profit += max - min
	return profit
}