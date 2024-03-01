package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(lengthOfLongestSubstring("abcabcbb")) // 3
	helper.PrintJSON(lengthOfLongestSubstring("bbbbb")) // 1
	helper.PrintJSON(lengthOfLongestSubstring("pwwkew")) // 3
	helper.PrintJSON(lengthOfLongestSubstring(" ")) // 1
	helper.PrintJSON(lengthOfLongestSubstring("")) // 0
}

/**
 * 滑动窗口。用 left 和 right 表示边界。right 往右移动，每次检查是否重复。
 **/
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	maxSubLen := 0
	left := 0
	right := -1
	for right < length - 1 {
		right++
		left = check(&s, left, right)
		if right - left + 1 > maxSubLen {
			maxSubLen = right - left + 1
		}
	}
	return maxSubLen
}

func check(s *string, left int, right int) int {
	for i := left; i < right; i++ {
		if (*s)[i] == (*s)[right] {
			return i + 1
		}
	}
	return left
}