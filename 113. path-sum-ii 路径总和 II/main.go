package main

import (
	"leetcode/go-helper"
)

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

type ListNode = helper.ListNode

func main() {
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	path := []int{}

	if root == nil {
		return res
	}

	var help func(root *TreeNode, targetSum int)
	help = func(root *TreeNode, targetSum int) {
		if root.Left == nil && root.Right == nil {
			if root.Val == targetSum {
				path = append(path, root.Val)
				res = append(res, append([]int{}, path...))
				path = path[ : len(path) - 1]
			}
		} else {
			if root.Left != nil {
				path = append(path, root.Val)
				help(root.Left, targetSum - root.Val)
				path = path[ : len(path) - 1]
			}
			if root.Right != nil {
				path = append(path, root.Val)
				help(root.Right, targetSum - root.Val)
				path = path[ : len(path) - 1]
			}
		}
	}

	help(root, targetSum)

	return res
}