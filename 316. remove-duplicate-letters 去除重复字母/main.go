package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(removeDuplicateLetters("bcabc")) // "abc"
	helper.PrintJSON(removeDuplicateLetters("cbacdcbc")) // "acdb"
	helper.PrintJSON(removeDuplicateLetters("leetcode")) // "letcod"
}

func removeDuplicateLetters(s string) string {
	// 字符统计
	charStatistic := make([]int, 26)
	for _, val := range s {
		idx := val - 'a'
		charStatistic[idx]++
	}
	// 字符是否在栈中
	inStack := make([]int, 26)
	stack := []rune{}
	// 遍历字符串
	for _, val := range s {
		idx := val - 'a'
		charStatistic[idx]--
		if inStack[idx] == 1 {
			continue
		}
		for len(stack) > 0 && stack[len(stack) - 1] > val && charStatistic[stack[len(stack) - 1] - 'a'] > 0 {
			inStack[stack[len(stack) - 1] - 'a'] = 0
			stack = stack[: len(stack) - 1]
		}
		stack = append(stack, val)
		inStack[idx] = 1
	}
	return string(stack)
}
