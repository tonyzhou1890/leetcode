package main

import (
	"leetcode/go-helper"
	"strconv"
)

func main() {
	helper.PrintJSON(countAndSay(1)) // "1"
	helper.PrintJSON(countAndSay(4)) // "1211"
	helper.PrintJSON(countAndSay(5)) // "111221"
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	res := ""
	count := 0
	var cur byte
	for i := 0; i < len(s); i++ {
		if s[i] == cur {
			count++
		} else {
			if count > 0 {
				res += strconv.Itoa(count) + string(cur)
			}
			count = 1
			cur = s[i]
		}
	}
	if count > 0 {
		res += strconv.Itoa(count) + string(cur)
	}
	return res
}