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
	helper.PrintJSON(generateTrees(3)) // [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
	helper.PrintJSON(generateTrees(1)) // [[1]]
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {

	var help func(start int, end int) []*TreeNode
	help = func(start, end int) []*TreeNode {
		temp := []*TreeNode{}

		for i := start; i <= end; i++ {
			// 左子树
			leftNodes := []*TreeNode{}
			leftNodes = append(leftNodes, nil)
			if start < i {
				leftNodes = help(start, i - 1)
			}
			// 右子树
			rightNodes := []*TreeNode{}
			rightNodes = append(rightNodes, nil)
			if i < end {
				rightNodes = help(i + 1, end)
			}
			// 组合
			for _, leftNode := range leftNodes {
				for _, rightNode := range rightNodes {
					node := &TreeNode{}
					node.Val = i
					node.Left = leftNode
					node.Right = rightNode
					temp = append(temp, node)
				}
			}
		}

		return temp
	}

	return help(1, n)
}
