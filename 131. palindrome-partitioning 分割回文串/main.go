package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(partition("aab")) // [["a","a","b"],["aa","b"]]
	helper.PrintJSON(partition("a")) // [["a"]]
}

func partition(s string) [][]string {
	length := len(s)

	res := [][]string{}
	path := []string{}

	// 是否回文字串
	isPalindrome := func(start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}
	// 递归分割
	var help func(start int)
	help = func(start int) {
		if start >= length {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := start; i < length; i++ {
			if isPalindrome(start, i) {
				path = append(path, s[start : i + 1])
				help(i + 1)
				path = path[ : len(path) - 1]
			}
		}
	}

	help(0)

	return res
}