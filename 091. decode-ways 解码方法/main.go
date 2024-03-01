package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(numDecodings("12")) // 2 它可以解码为 "AB"（1 2）或者 "L"（12）。
	helper.PrintJSON(numDecodings("226")) // 3 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
	helper.PrintJSON(numDecodings("06")) // 0 "06" 无法映射到 "F" ，因为存在前导零（"6" 和 "06" 并不等价）。
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] != '0' {
				dp[i] = 1
			}
		}
		if i > 0 {
			// s[i] 单独
			if s[i] != '0' {
				dp[i] = dp[i - 1]
			}
			// s[i] 与前一个字符组合
			if (s[i - 1] == '1') || (s[i - 1] == '2' && s[i] < '7') {
				if i > 1 {
					dp[i] += dp[i - 2]
				} else {
					dp[i] += 1
				}
			}
		}
	}
	return dp[len(dp) - 1]
}