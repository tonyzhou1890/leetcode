package main

import (
	"leetcode/go-helper"
)

func main() {
	// helper.PrintJSON(strStr("sadbutsad", "sad")) // 0
	// helper.PrintJSON(strStr("leetcode", "leeto")) // -1
	// helper.PrintJSON(strStr("leetcode", "code")) // 4
	// helper.PrintJSON(strStr("babbbbbabb", "bbab")) // 5
	helper.PrintJSON(strStr("mississippi", "pi")) // 9
}

func strStr(haystack string, needle string) int {
	res := -1
	if len(haystack) < len(needle) {
		return res
	}
	for i, j := 0, 0; i < len(haystack) && j < len(needle); {
		// 最后一个字符匹配
		if j == len(needle) - 1 && haystack[i] == needle[j] {
			res = i - j
			break
		}
		// 当遇到不同字符，在 needle 中 j 之前查找此不同字符，找到，则对齐字符，从头开始匹配；找不到，则 needle 整体移动到此字符后面
		if haystack[i] != needle[j] {
			k := j - 1
			flag := false
			for k >= 0 {
				if haystack[i] == needle[k] {
					flag = true
					break
				}
				k--
			}
			if flag {
				i -= k
				j = 0
				continue
			} else {
				j = 0
				i++
				continue
			}
		}
		i++
		j++
	}
	return res
}
