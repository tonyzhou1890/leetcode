package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(maxProfit([]int{1,2,3,0,2})) // 返回 3
	helper.PrintJSON(maxProfit([]int{1})) // 返回 0
	helper.PrintJSON(maxProfit([]int{2,1,4})) // 3
}

/**
 * 维护两个数组，一个代表第 n 天持有的最大收益，一个代表第 n 天未持有的最大收益
 * 第 n 天持有的最大收益可以从前一天继续持有以及 n-2 天未持有购买转化而来
 * 第 n 天未持有的最大收益可以从前一天继续未持有以及 n-1 天持有卖出转化而来
 **/
func maxProfit(prices []int) int {
	length := len(prices)
	var has = make([]int, length)
	var not = make([]int, length)
	has[0] = 0 - prices[0]
	for i := 1; i < length; i++ {
		// 持有
		if i == 1 {
			if 0 - prices[i] > has[0] {
				has[i] = 0 - prices[i]
			} else {
				has[i] = has[0]
			}
		} else {
			if not[i - 2] - prices[i] > has[i - 1] {
				has[i] = not[i - 2] - prices[i]
			} else {
				has[i] = has[i - 1]
			}
		}
		// 未持有
		if has[i - 1] + prices[i] > not[i - 1] {
			not[i] = has[i - 1] + prices[i]
		} else {
			not[i] = not[i - 1]
		}
	}
	if has[length - 1] > not[length - 1] {
		return has[length - 1]
	}
	return not[length - 1]
}
