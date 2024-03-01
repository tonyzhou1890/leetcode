package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(maxProfit([]int{7,1,5,3,6,4})) // 5
	helper.PrintJSON(maxProfit([]int{7,6,4,3,1})) // 0
}


func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] - min > profit {
			profit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}