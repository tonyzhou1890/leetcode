package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(isInterleave("aabcc", "dbbca", "aadbbcbcac")) // true
	helper.PrintJSON(isInterleave("aabcc", "dbbca", "aadbbbaccc")) // false
	helper.PrintJSON(isInterleave("", "", "")) // true
	helper.PrintJSON(isInterleave("", "", "a")) // false
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)

	if m + n != len(s3) {
		return false
	}

	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n + 1)
	}

	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i > 0 {
				if s1[i - 1] == s3[i + j - 1] {
					dp[i][j] = dp[i - 1][j]
				}
			}
			if j > 0 {
				if s2[j - 1] == s3[i + j - 1] && dp[i][j] == false {
					dp[i][j] = dp[i][j - 1]
				}
			}
		}
	}

	return dp[m][n]
}