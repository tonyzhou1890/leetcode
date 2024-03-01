package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(mySqrt(4)) // 2
	helper.PrintJSON(mySqrt(8)) // 2
	helper.PrintJSON(mySqrt(0)) // 0
	helper.PrintJSON(mySqrt(1)) // 1
	helper.PrintJSON(mySqrt(2)) // 1
	helper.PrintJSON(mySqrt(3)) // 1
	helper.PrintJSON(mySqrt(5)) // 2
}

// func mySqrt(x int) int {
// 	if x == 0 || x == 1 {
// 		return x
// 	}
// 	l := 0
// 	r := x
// 	for l <= r {
// 		m := (l + r) / 2
// 		if x / m < m {
// 			r = m - 1
// 		} else {
// 			l = m + 1
// 		}
// 	}
// 	return l - 1
// }

// 牛顿迭代法
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	c := float64(x)
	x0 := c
	for ;; {
		xi := 0.5 * (x0 + c / x0)
		if x0 - xi <= 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}