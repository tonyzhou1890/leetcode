package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(myPow(2.00000, 10)) // 1024.00000
	helper.PrintJSON(myPow(2.10000, 3)) // 9.26100
	helper.PrintJSON(myPow(2.00000, -2)) // 0.25000
	helper.PrintJSON(myPow(0.44528, 0)) // 1
}

func myPow(x float64, n int) float64 {
	// 负幂处理
	if n < 0 {
		x  = 1 / x
		// 防止溢出
		if n == - (2 << 31) {
			n /= 2
			x *= x
		}
		n = -n
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n % 2 == 1 {
		return myPow(x * x, n >> 1) * x
	} else {
		return myPow(x * x, n >> 1)
	}
}