package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(wordBreak("leetcode", []string{"leet", "code"})) // true
	helper.PrintJSON(wordBreak("applepenapple", []string{"apple", "pen"})) // true
	helper.PrintJSON(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}

// 动态规划，dp[i] 表示 0 -- i-1 的字符能不能由 wordDict 组成
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s) + 1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				temp := s[j : i]
				for k := 0; k < len(wordDict); k++ {
					if temp == wordDict[k] {
						dp[i] = true
						break
					}
				}
			}
		}
	}
	return dp[len(s)]
}