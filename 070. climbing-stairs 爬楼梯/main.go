package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(climbStairs(2)) // 2
	helper.PrintJSON(climbStairs(3)) // 3
	helper.PrintJSON(climbStairs(4)) // 5
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n - 1]
}