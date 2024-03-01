package main

import (
	"leetcode/go-helper"
	"strconv"
)

func main() {
	helper.PrintJSON(addBinary("11", "1")) // "100"
	helper.PrintJSON(addBinary("1010", "1011")) // "10101"
}

func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	if lenA < lenB {
		a, b = b, a
		lenA, lenB = lenB, lenA
	}
	carry := 0
	cur := 0
	res := ""
	for i := lenA - 1; i >= 0; i-- {
		cur = int((a[i] - '0')) + carry
		if i >= lenA - lenB {
			cur += int((b[i - lenA + lenB] - '0'))
		}
		carry = cur / 2
		cur %= 2
		res = strconv.Itoa(cur) + res
	}
	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}