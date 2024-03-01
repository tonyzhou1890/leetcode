package main

import (
	"leetcode/go-helper"
	"math"
)

func main() {
	helper.PrintJSON(myAtoi("42")) // 42
	helper.PrintJSON(myAtoi("   -42")) // -42
	helper.PrintJSON(myAtoi("4193 with words")) // 4193
	helper.PrintJSON(myAtoi("00-1")) // 0
	helper.PrintJSON(myAtoi("+1")) // 1
	helper.PrintJSON(myAtoi("21877637271232")) // 2187763727
	helper.PrintJSON(myAtoi("-21877637271232")) // -2147483648
}

func myAtoi(s string) int {
	intMax := int(math.Pow(2, 31) - 1)
	intMin := int(-math.Pow(2, 31))
	res := 0
	flag := 0 // 1 正，-1 负
	for _, val := range s {
		// 前导空格
		if val == ' ' {
			if flag == 0 {
				continue
			} else {
				break
			}
		}
		// 数字
		if val >= '0' && val <= '9' {
			if flag == 0 {
				flag = 1
			}
			res = res * 10 + int(val - '0')
			if res * flag >= intMax {
				return intMax
			}
			if res * flag <= intMin {
				return intMin
			}
		} else if val == '-' && flag == 0 { // 符号--当前没有符号并且无值才有效
			flag = -1
		} else if val == '+' && flag == 0 {
			flag = 1
		} else {
			break
		}
	}
	return res * flag
}
