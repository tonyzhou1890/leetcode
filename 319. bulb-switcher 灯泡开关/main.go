package main

import (
	"leetcode/go-helper"
	"math"
)

func main() {
	helper.PrintJSON(bulbSwitch(3)) // 1
	helper.PrintJSON(bulbSwitch(0)) // 0
	helper.PrintJSON(bulbSwitch(1)) // 1
}

/**
 * 灯泡是否亮取决于因数数量。平方数的约数是奇数个（比如：4 = 2 * 2 = 1 * 4），非平方数的约数是偶数个。
 * 所以，问题变成 n 以内平方数的个数。因此，答案是 sqrt(n)
 **/
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
