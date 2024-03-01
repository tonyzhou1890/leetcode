package main

import (
	"leetcode/go-helper"
	"math"
)

func main() {
	helper.PrintJSON(divide(10, 3)) // 3
	helper.PrintJSON(divide(7, -3)) // -2
	helper.PrintJSON(divide(37228472, -3)) // -12409490
	helper.PrintJSON(divide(-1 << 31, -1)) // 1 << 31 - 1
}

func divide(dividend int, divisor int) int {
	if dividend == -1 << 31 && divisor == -1 {
		return 1 << 31 - 1
	}
	if dividend == 0 {
		return 0
	}
	flag := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		flag = -1
	}
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))
	res := 0
	for m := 31; m >= 0; m-- {
		if (dividend >> m) >= divisor {
			res += 1 << m
			dividend -= divisor << m
			m = 31
		}
	}
	if flag > 0 {
		return res
	} else {
		return -res
	}
}