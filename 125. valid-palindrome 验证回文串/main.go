package main

import (
	"leetcode/go-helper"
	"regexp"
	"strings"
)

func main() {
	helper.PrintJSON(isPalindrome("A man, a plan, a canal: Panama")) // true
	helper.PrintJSON(isPalindrome("race a car")) // false
	helper.PrintJSON(isPalindrome(" ")) // true
}


// func isPalindrome(s string) bool {
// 	left := 0
// 	right := len(s) - 1

// 	s = strings.ToLower(s)
	
// 	for left < right {
// 		for left < right && !((s[left] >= '0' && s[left] <= '9') || (s[left] >= 'a' && s[left] <= 'z')) {
// 			left++
// 		}
// 		for left < right && !((s[right] >= '0' && s[right] <= '9') || (s[right] >= 'a' && s[right] <= 'z')) {
// 			right--
// 		}
// 		if left >= right {
// 			return true
// 		}
// 		if s[left] != s[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

func isPalindrome(s string) bool {
	reg := regexp.MustCompile("[^0-9a-zA-Z]")
	s = reg.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	left := 0
	right := len(s) - 1
	
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}