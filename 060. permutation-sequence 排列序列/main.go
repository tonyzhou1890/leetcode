package main

import (
	"leetcode/go-helper"
	"strconv"
	"strings"
)

func main() {
	helper.PrintJSON(getPermutation(3, 3)) // "213"
	helper.PrintJSON(getPermutation(4, 9)) // "2314"
	helper.PrintJSON(getPermutation(3, 1)) // "123"
}

// // 通过字典序生成排列，返回第 k 个
// func getPermutation(n int, k int) string {
// 	nums := make([]int, n)
// 	res := make([]string, n)
// 	for i := 0; i < n; i++ {
// 		nums[i] = i + 1
// 	}
// 	// i 从 1 开始，因为 nums 本身就是第一个排列，只需要生成 k - 1 个排列就可以了
// 	for i := 1; i < k; i++ {
// 		nums, _ = next(nums)
// 	}
// 	for i := 0; i < n; i++ {
// 		res[i] = strconv.Itoa(nums[i])
// 	}
// 	return strings.Join(res, "")
// }

// // 下一个排列。返回值：下一个排列，是否回到第一个排列
// func next(nums []int) ([]int, bool) {
// 	l := -1
// 	r := 0
// 	isFirst := false
// 	// 找峰值左侧索引
// 	for i := len(nums) - 1; i > 0; i-- {
// 		if nums[i] > nums[i - 1] {
// 			l = i - 1
// 			// 找峰值（包含）右侧最后一个比 l 处大的值
// 			for j := len(nums) - 1; j > l; j-- {
// 				if nums[j] > nums[l] {
// 					r = j
// 					break
// 				}
// 			}
// 			nums[l], nums[r] = nums[r], nums[l]
// 			break
// 		}
// 	}
// 	// 峰值右侧变升序
// 	for i, j := l + 1, len(nums) - 1; i < j; i, j = i + 1, j - 1 {
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// 	if l == -1 {
// 		isFirst = true
// 	}
// 	return nums, isFirst
// }

/*
1, 2, 3, 4
01. [1,2,3,4],
02. [1,2,4,3],
03. [1,3,2,4],
04. [1,3,4,2],
05. [1,4,2,3],
06. [1,4,3,2],
07. [2,1,3,4],
08. [2,1,4,3],
09. [2,3,1,4],
10. [2,3,4,1],
11. [2,4,1,3],
12. [2,4,3,1],
13. [3,1,2,4],
14. [3,1,4,2],
15. [3,2,1,4],
16. [3,2,4,1],
17. [3,4,1,2],
18. [3,4,2,1],
19. [4,1,2,3],
20. [4,1,3,2],
21. [4,2,1,3],
22. [4,2,3,1],
23. [4,3,1,2],
24. [4,3,2,1]

res[m] = nums[k / (n - 1)!]
*/

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	// 序列从 0 开始计数，k 要减去 1
	k -= 1
	for i := 0; i < n; i++ {
		f := fac(n - 1 - i)
		res[i] = strconv.Itoa(nums[k / f])
		nums = append(nums[ : k / f], nums[k / f + 1 : ]...)
		k %= f
	}
	return strings.Join(res, "")
}

func fac(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}