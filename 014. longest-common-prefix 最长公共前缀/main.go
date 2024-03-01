package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(longestCommonPrefix([]string{"flower","flow","flight"})) // "fl"
	helper.PrintJSON(longestCommonPrefix([]string{"dog","racecar","car"})) // ""
	helper.PrintJSON(longestCommonPrefix([]string{"IX"})) // "IX"
}

func longestCommonPrefix(strs []string) string {
	min := 2 << 32 - 1
	for i := 0; i < len(strs); i++ {
		sLen := len(strs[i])
		if sLen < min {
			min = sLen
		}
	}
	for i := 0; i < min; i++ {
		temp := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if temp != strs[j][i] {
				return strs[j][0:i]
			}
		}
	}
	return strs[0][0:min]
}