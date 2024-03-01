package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(lengthOfLastWord("Hello World")) // 5
	helper.PrintJSON(lengthOfLastWord("   fly me   to   the moon  ")) // 4
	helper.PrintJSON(lengthOfLastWord("luffy is still joyboy")) // 6
}

// func lengthOfLastWord(s string) int {
// 	end := 0
// 	start := 0
// 	for i := len(s) - 1; i >= 0; i-- {
// 		// 遇到第一个字母
// 		if end == 0 && ((s[i] >= 65 && s[i] < (65 + 26)) || (s[i] >= 97 && s[i] < (97 + 26))) {
// 			end = i
// 			continue
// 		}
// 		// 遇到第一个空格
// 		if end != 0 && s[i] == ' ' {
// 			start = i + 1
// 			break
// 		}
// 	}
// 	return end - start + 1
// }

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if length == 0 && s[i] == ' ' {
			continue
		}
		if length != 0 && s[i] == ' ' {
			break
		}
		length++
	}
	return length
}