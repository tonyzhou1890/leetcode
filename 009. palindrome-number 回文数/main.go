package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(isPalindrome(121)) // true
	helper.PrintJSON(isPalindrome(-121)) // false
	helper.PrintJSON(isPalindrome(10)) // false
	helper.PrintJSON(isPalindrome(0)) // true
}

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || (x % 10) == 0 {
		return false
	}
	raw := x
	reverseX := 0
	for raw > 0 {
		reverseX = reverseX * 10 + (raw % 10)
		raw /= 10
	}
	return reverseX == x
}