package main

import (
	"leetcode/go-helper"
	"strings"
)

func main() {
	helper.PrintJSON(reverseWords("the sky is blue")) // "blue is sky the"
	helper.PrintJSON(reverseWords("  hello world  ")) // "world hello"
	helper.PrintJSON(reverseWords("a good   example")) // "example good a"
}

func reverseWords(s string) string {
	words := []string{}
	word := ""

	for i := len(s) - 1; i >= 0; i-- {
		if word != "" && s[i] == ' ' {
			words = append(words, word)
			word = ""
		} else if s[i] != ' ' {
			word = string(s[i]) + word
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return strings.Join(words, " ")
}