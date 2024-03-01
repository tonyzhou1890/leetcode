package main

import (
	"leetcode/go-helper"
)

func main() {
	helper.PrintJSON(grayCode(1)) // [0,1]
	helper.PrintJSON(grayCode(2)) // [0,1,3,2]
	helper.PrintJSON(grayCode(3)) // [0,1,3,2]
	helper.PrintJSON(grayCode(4)) // [0,1,3,2]
}

// n 位格雷码前 2^(n-1) 个结果为 n-1 位格雷码最前面加 0。后 2^(n-1) 个结果为前面一半的逆序，然后最前面一位变成 1
// func grayCode(n int) []int {
// 	res := []int{}
// 	if n == 1 {
// 		res = append(res, 0)
// 		res = append(res, 1)
// 		return res
// 	}
// 	res = grayCode(n - 1)
// 	length := len(res)
// 	for i := length - 1; i >= 0; i-- {
// 		res = append(res, res[i] | (1 << (n - 1)))
// 	}
// 	return res
// }

// 直接由 2 进制转格雷码。格雷码的第 i 位为二进制第 i + 1 位和第 i 位异或。
func grayCode(n int) []int {
	res := make([]int, 2 << (n - 1))
	for i := 0; i < len(res); i++ {
		// 这里 i 右移 1 位与 i 异或，就是把 i 的每一位与后面一位异或。
		res[i] = (i >> 1) ^ i
	}
	return res
}