package main

import (
	"leetcode/go-helper"
	"strconv"
)

func main() {
	helper.PrintJSON(multiply("2", "3")) // "6"
	helper.PrintJSON(multiply("123", "456")) // "56088"
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	length := len(num1) + len(num2)
	// 实际长度为 length 或者 length - 1。比如 1000 * 1，长度为 length - 1。9999 * 9，长度为 length
	res := make([]int, length)

	// 模拟每位相乘
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			res[i + j + 1] += int(num1[i] - '0') * int(num2[j] - '0')
		}
	}

	// 每位相加
	carry := 0
	for i := length - 1; i >= 0; i-- {
		res[i] += carry
		carry = res[i] / 10
		res[i] %= 10
	}

	start := 0
	if res[0] == 0 {
		start = 1
	}
	s := ""
	for start < length {
		s += strconv.Itoa(res[start])
		start++
	}
	return s
}