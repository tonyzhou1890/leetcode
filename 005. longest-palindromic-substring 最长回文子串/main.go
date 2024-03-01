package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(longestPalindrome("babad")) // bab aba
	helper.PrintJSON(longestPalindrome("cbbd")) // bb
	helper.PrintJSON(longestPalindrome("a")) // a
}

/**
 * 中心扩散
 **/
func longestPalindrome(s string) string {
	length := len(s)
	maxLength := 1
	left := 0
	right := 0
	for i := 0; i < length; i++ {
		l1, r1 := check(&s, length, i, i)
		l2, r2 := check(&s, length, i, i + 1)
		if r1 - l1 + 1 > maxLength {
			maxLength = r1 - l1 + 1
			left = l1
			right = r1
		}
		if l2 <= r2 && r2 - l2 + 1 > maxLength {
			maxLength = r2 - l2 + 1
			left = l2
			right = r2
		}
	}
	return s[left : right + 1]
}

func check(s *string, length int, left int, right int) (int, int) {
	for left >= 0 && right < length {
		if (*s)[left] != (*s)[right] {
			break
		}
		left--
		right++
	}
	return left + 1, right - 1
}