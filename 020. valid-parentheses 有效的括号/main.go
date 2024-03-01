package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(isValid("()")) // true
	helper.PrintJSON(isValid("()[]{}")) // true
	helper.PrintJSON(isValid("(]")) // false
}

func isValid(s string) bool {
	valid := true
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				valid = false
				break
			}
			last := stack[len(stack) - 1]
			if (last == '(' && c == ')') || (last == '[' && c == ']') || (last == '{' && c == '}') {
				stack = stack[:len(stack) - 1]
			} else {
				valid = false
				break
			}
		}
	}
	if len(stack) > 0 {
		valid = false
	}
	return valid
}
