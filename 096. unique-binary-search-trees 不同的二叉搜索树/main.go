package main

import (
	"leetcode/go-helper"
)

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

func main() {
	helper.PrintJSON(numTrees(3)) // 5
	helper.PrintJSON(numTrees(4)) // 14
	helper.PrintJSON(numTrees(5)) // 42
	helper.PrintJSON(numTrees(6)) // 132
	helper.PrintJSON(numTrees(7)) // 429
	helper.PrintJSON(numTrees(18)) // 477638700
	helper.PrintJSON(numTrees(1)) // 1
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func numTrees(n int) int {

// 	cache := make([][]int, n + 1)
// 	for i := 0; i < n + 1; i++ {
// 		cache[i] = make([]int, n + 1)
// 	}

// 	var help func(start int, end int) int
// 	help = func(start, end int) int {
// 		if cache[start][end] != 0 {
// 			return cache[start][end]
// 		}
// 		count := 0
// 		for i := start; i <= end; i++ {
// 			// 左子树种类
// 			leftCount := 1
// 			if start < i {
// 				leftCount = help(start, i - 1)
// 			}
// 			// 右子树种类
// 			rightCount := 1
// 			if end > i {
// 				rightCount = help(i + 1, end)
// 			}
// 			count += leftCount * rightCount
// 		}
// 		cache[start][end] = count
// 		return count
// 	}

// 	return help(1, n)
// }

/**
G(n) 表示 n 个数的搜索树数量
F(k, n) 表示 n 个数以第 k 个数为根节点的搜索树数量
比如：1,2,3,4 和 5,6,7,8 组成的搜索树数量是一样的，都是 G(4)
因而，F(k, n) = G(k - 1) * G(n - k)
*/
func numTrees(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j - 1] * dp[i - j]
		}
	}
	return dp[n]
}