package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(maxProduct([]string{"abcw","baz","foo","bar","xtfn","abcdef"})) // 16
	helper.PrintJSON(maxProduct([]string{"a","ab","abc","d","cd","bcd","abcd"})) // 4
	helper.PrintJSON(maxProduct([]string{"a","aa","aaa","aaaa"})) // 0
}

/**
 * 遍历处理。
 * 先比较字符串。用 32 位二进制表示 26 个字母是否存在。然后两个二进制数按位与。结果为 0，则计算长度乘积。
 **/
func maxProduct(words []string) int {
	maxLength := 0
	masks := make([]int, len(words))
	for idx, val := range words {
		for _, char := range val {
			masks[idx] |= 1 << (char - 'a')
		}
	}
	for i := 0; i < len(words) - 1; i++ {
		for j := i + 1; j < len(words); j++ {
			if (masks[i] & masks[j]) > 0 {
				continue
			}
			product := len(words[i]) * len(words[j])
			if product > maxLength {
				maxLength = product
			}
		}
	}
	return maxLength
}
