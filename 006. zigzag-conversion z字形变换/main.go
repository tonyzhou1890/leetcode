package main

import (
	"leetcode/go-helper"
	"strings"
)

func main() {
	helper.PrintJSON(convert("PAYPALISHIRING", 3)) // "PAHNAPLSIIGYIR"
	helper.PrintJSON(convert("PAYPALISHIRING", 4)) // "PINALSIGYAHRPI"
	helper.PrintJSON(convert("A", 1)) // "A"
	helper.PrintJSON(convert("AB", 1)) // "AB"
}

/**
 * 上下游走模拟
 **/
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	// 第几行
	i := 0
	// flag 为 1，向下，-1 向上
	flag := -1
	res := make([]string, numRows)
	for _, c := range s {
		res[i] += string(c)
		if i == 0 || i == numRows -1 {
			flag = -flag
		}
		i += flag
	}
	return strings.Join(res, "")
}