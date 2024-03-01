package main

import (
	"leetcode/go-helper"
	"math"
)

func main() {
	helper.PrintJSON(reverse(123)) // 321
	helper.PrintJSON(reverse(-123)) // -321
	helper.PrintJSON(reverse(120)) // 21
	helper.PrintJSON(reverse(0)) // 0
	helper.PrintJSON(reverse(int(math.Pow(2, 31) - 1))) // 0
	helper.PrintJSON(reverse(int(-math.Pow(2, 31)))) // 0
	helper.PrintJSON(reverse(int(math.Pow(2, 31) - 46))) // 0
	helper.PrintJSON(reverse(int(-math.Pow(2, 31) + 7))) // 0
}

func reverse(x int) int {
	nineMax := int(math.Pow(2, 31) - 1)
	nineMin := int(-math.Pow(2, 31))
	r := 0
	for x != 0 {
		r = r * 10 + (x % 10)
		if r > nineMax || r < nineMin {
			return 0
		}
		x /= 10
	}
	return r
}